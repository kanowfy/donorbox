import { Avatar, Button, Modal, Banner, Tooltip } from "flowbite-react";
import { useParams, useNavigate, Link } from "react-router-dom";
import Support from "../../components/Support";
import DonateBox from "../../components/DonateBox";
import { useEffect, useState } from "react";
import projectService from "../../services/project";
import utils from "../../utils/utils";
import { IoFlag } from "react-icons/io5";
import MilestoneTL from "../../components/MilestoneTL";
import MDEditor from "@uiw/react-md-editor";
import { HiX } from "react-icons/hi";
import { IoReceipt } from "react-icons/io5";

const Project = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [project, setProject] = useState({});
  const [milestones, setMilestones] = useState([]);
  const [owner, setOwner] = useState({});
  const [backings, setBackings] = useState();
  const [proofs, setProofs] = useState();
  const [wosList, setWosList] = useState();
  const [isOpenMilestone, setIsOpenMilestone] = useState(false);
  const [milestoneReview, setMilestoneReview] = useState();
  const [displayImages, setDisplayImages] = useState();

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const projectResponse = await projectService.getOne(params.id);
        console.log(projectResponse);
        setProject(projectResponse.project);
        setMilestones(projectResponse.milestones);
        setBackings(projectResponse.backings);
        setOwner(projectResponse.user);

        setProofs(
          projectResponse.milestones
            .filter((m) => m?.spending_proofs?.length > 0)
            .flatMap((m) =>
              m.spending_proofs.map((p) => ({
                ...p,
                milestone_title: m.title,
              }))
            )
            .filter((p) => p.status === "approved")
        );

        if (projectResponse.backings) {
          setWosList(
            projectResponse.backings.filter(
              (b) => b.word_of_support !== undefined
            )
          );
        }
      } catch (err) {
        navigate("/not-found");
        console.error(err);
      }
    };

    fetchProject();
  }, [params.id, navigate]);

  return (
    <div className="mb-10">
      <StatusBanner status={project?.status} />
      <div className="grid grid-cols-3 gap-7 max-w-7xl min-w-1/2 mx-auto">
        <div className="col-span-2">
          <div className="text-3xl font-bold m-5">{project.title}</div>
          <div className="rounded-xl overflow-hidden min-h-80 max-h-128 aspect-[5/3] object-cover">
            <img
              src={project.cover_picture}
              className="w-full h-full m-auto object-cover"
            />
          </div>
          <div className="m-3 flex justify-between">
            <Avatar
              alt="User settings"
              img={
                owner.profile_picture ? owner.profile_picture : "/avatar.svg"
              }
              rounded
            >
              <div>
                <span className="flex justify-start">
                  {`${owner.first_name} ${owner.last_name}`} is organizing this
                  fundraiser.
                </span>
                <span className="block font-normal text-gray-500 text-sm">
                  Created{" "}
                  {utils.calculateDayDifference(
                    Date.now(),
                    utils.parseDateFromRFC3339(project.created_at)
                  )}
                  d ago.
                </span>
              </div>
            </Avatar>
            <div className="flex flex-col items-end">
              <div className="font-medium">
                {project.city}, {project.country}
              </div>
              <div className="flex space-x-1 items-end text-gray-500 text-sm">
                <div>Fundraiser ends on </div>
                <div className="font-medium">
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(project.end_date))
                  )}
                </div>
              </div>
            </div>
          </div>
          <div className="h-px bg-gray-300 mb-4"></div>

          <div className="text-xl font-semibold tracking-tight mb-3">
            About the Fundraiser
          </div>
          {/*<p className="max-w-2xl tracking-tight">{project.description}</p>*/}
          <div data-color-mode="light">
            <MDEditor.Markdown
              source={project.description}
              style={{ whiteSpace: "pre-wrap" }}
            />
          </div>

          <div className="h-px bg-gray-300 mt-4"></div>

          {proofs?.length > 0 && (
            <>
              <div className="my-5">
                <div className="text-xl font-semibold tracking-tight mb-5">
                  Proof of expenditure ({proofs.length})
                </div>
                <div className="space-y-4">
                  {proofs?.map((p) => (
                    <div key={p.id} className="p-5 rounded-lg border space-y-3">
                      <div className="flex justify-between">
                        <div>Milestone: <span className="font-semibold">{p.milestone_title}</span></div>
                        <div className="font-medium text-sm text-gray-600">
                          {utils.formatDate(
                            new Date(utils.parseDateFromRFC3339(p.created_at))
                          )}
                        </div>
                      </div>

                      <Button size="sm" color="light" onClick={() => setDisplayImages(p.transfer_image)}>
                        <IoReceipt className="w-5 h-5 mr-2" />
                        View transfer receipts
                      </Button>
                      <div className="rounded-lg border border-gray-300 overflow-hidden h-96 object-cover my-2 mx-auto" onClick={() => setDisplayImages(p.proof_media)}>
                        <img
                          src={p.proof_media}
                          className="w-full h-full m-auto object-cover filter hover:brightness-75 cursor-pointer"
                        />
                      </div>
                      <div className="tracking-tight"><span>Note: </span>{p.description}</div>
                    </div>
                  ))}
                </div>
              </div>
              <div className="h-px bg-gray-300"></div>
            </>
          )}

          <div className="my-5">
            <div className="text-xl font-semibold tracking-tight">
              Donors&apos; words of support (
              {wosList?.length ? wosList.length : 0})
            </div>

            {wosList?.map((b) => (
              <Support
                key={b.id}
                avatar={
                  b.backer.profile_picture
                    ? b.backer.profile_picture
                    : "/avatar.svg"
                }
                first_name={
                  b.backer.first_name ? b.backer.first_name : "Anonymous"
                }
                last_name={b.backer.last_name ? b.backer.last_name : ""}
                amount={b.amount}
                day_since={utils.getDaySince(b.created_at)}
                comment={b.word_of_support}
              />
            ))}
          </div>
          <div className="h-px bg-gray-300"></div>
          <div className="my-5">
            <Link to={`/fundraiser/${params.id}/report`} className="w-fit">
              <Button color="light" pill size="lg">
                <IoFlag className="w-5 h-5 mr-1" />
                Report Fundraiser
              </Button>
            </Link>
          </div>
        </div>
        <div className="mt-20 space-y-10">
          <div>
            <DonateBox
              id={params.id}
              totalFund={project.total_fund}
              fundGoal={project.fund_goal}
              backings={backings}
              status={project.status}
            />
          </div>
          <div className="p-4">
            <div className="text-2xl mb-4">Milestones</div>
            <MilestoneTL
              milestones={milestones}
              setIsOpenMilestone={setIsOpenMilestone}
              setMilestoneReview={setMilestoneReview}
            />
          </div>
        </div>
      </div>
      <Modal
        show={isOpenMilestone}
        onClose={() => setIsOpenMilestone(false)}
        size="3xl"
      >
        <Modal.Header>Milestone Fund Release Receipt</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col items-center justify-center space-y-2">
            <div className="text-yellow-500 text-5xl">{`₫${milestoneReview?.milestone_completion.transfer_amount.toLocaleString()}`}</div>
            <div className="text-gray-600">transferred on</div>
            <div className="text-green-600 text-2xl">
              {utils.formatDate(
                new Date(
                  utils.parseDateFromRFC3339(
                    milestoneReview?.milestone_completion.created_at
                  )
                )
              )}
            </div>
          </div>
          <div className="flex justify-center my-7">
          <div className="space-y-3">
          {milestoneReview?.milestone_completion.transfer_image && utils.parseImageUrl(milestoneReview?.milestone_completion.transfer_image).map(i => (
            <img src={i} alt={i}/>
          ))}
          </div>
          </div>
          {milestoneReview?.milestone_completion.transfer_note && (
            <div className="flex space-x-1 mx-4">
              <div className="text-gray-900">
                Note: {milestoneReview?.milestone_completion.transfer_note}
              </div>
            </div>
          )}
        </Modal.Body>
      </Modal>
      <Modal
        show={displayImages != null}
        onClose={() => setDisplayImages(null)}
        size="7xl"
        popup
      >
        <Modal.Header />
        <Modal.Body>
          <div className="space-y-3">
          {displayImages && utils.parseImageUrl(displayImages).map(i => (
            <img src={i} alt={i}/>
          ))}

          </div>
        </Modal.Body>

      </Modal>
    </div>
  );
};

const StatusBanner = ({ status }) => {
  const statusList = {
    pending: {
      text: "This fundraiser is currently under review",
      color: "text-green-600",
    },
    disputed: {
      text: "This fundraiser is under investigation due to violation of the platform policy",
      color: "text-yellow-600",
    },
    stopped: {
      text: "This fundraiser is archived due to violation of platform policy",
      color: "text-red-700",
    },
    rejected: {
      text: "This fundraiser is not eligible for launch",
      color: "text-yellow-600",
    },
  };

  if (!statusList.hasOwnProperty(status)) {
    return <></>;
  }

  return (
    <Banner>
      <div className="flex w-full justify-between border-b border-gray-200 bg-gray-50 p-4">
        <div className="mx-auto flex items-center">
          <p
            className={`flex items-center text-sm font-semibold ${statusList[status].color}`}
          >
            {statusList[status].text}
          </p>
        </div>
        <Banner.CollapseButton
          color="gray"
          className="border-0 bg-transparent text-gray-500 dark:text-gray-400"
        >
          <HiX className="h-4 w-4" />
        </Banner.CollapseButton>
      </div>
    </Banner>
  );
};

export default Project;
