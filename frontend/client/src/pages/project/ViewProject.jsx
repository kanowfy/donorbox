import { Avatar, Button } from "flowbite-react";
import { useParams, useNavigate, Link } from "react-router-dom";
import Support from "../../components/Support";
import DonateBox from "../../components/DonateBox";
import { useEffect, useState } from "react";
import projectService from "../../services/project";
import utils from "../../utils/utils";
import { IoFlag } from "react-icons/io5";
import MilestoneTL from "../../components/MilestoneTL";

const Project = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [project, setProject] = useState({});
  const [milestones, setMilestones] = useState([]);
  const [owner, setOwner] = useState({});
  const [backings, setBackings] = useState();
  const [updates, setUpdates] = useState();
  const [wosList, setWosList] = useState();

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const projectResponse = await projectService.getOne(params.id);
        console.log(projectResponse);
        setProject(projectResponse.project);
        setMilestones(projectResponse.milestones);
        setBackings(projectResponse.backings);
        setUpdates(projectResponse.updates);
        setOwner(projectResponse.user);

        if (projectResponse.backings) {
          setWosList(
            projectResponse.backings.filter((b) => b.word_of_support !== undefined)
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
          <p className="max-w-2xl tracking-tight">{project.description}</p>

          <div className="h-px bg-gray-300 mt-4"></div>

          {updates && (
            <>
              <div className="my-5">
                <div className="text-xl font-semibold tracking-tight mb-5">
                  Updates ({updates.length})
                </div>
                <div className="space-y-4">
                  {updates.map((u) => (
                    <div key={u.id}>
                      <div className="font-medium text-sm">
                        On{" "}
                        {utils.formatDate(
                          new Date(utils.parseDateFromRFC3339(u.created_at))
                        )}
                      </div>
                      <div className="tracking-tight">{u.description}</div>
                      {u?.attachment_photo && (
                        <div className="rounded-xl overflow-hidden h-40 aspect-[4/3] object-cover my-2">
                          <img
                            src={u.attachment_photo}
                            className="w-full h-full m-auto object-cover"
                          />
                        </div>
                      )}
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
                avatar={b.backer.profile_picture ? b.backer.profile_picture : "/avatar.svg"}
                first_name={b.backer.first_name ? b.backer.first_name : "Anonymous"}
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
          />
          </div>
          <div className="p-4">
            <div className="text-2xl mb-4">Milestones</div>
            <MilestoneTL milestones={milestones}/>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Project;
