import { Avatar } from "flowbite-react";
import { Link, useParams, useNavigate } from "react-router-dom";
import Support from "../../components/Support";
import DonateBox from "../../components/DonateBox";
import { useEffect, useState } from "react";
import projectService from "../../services/project";
import utils from "../../utils/utils";

const Project = () => {
  // useparams to get id
  // useeffect to fetch the project and all related shi
  // plug that shi in
  const params = useParams();
  const navigate = useNavigate();
  const [project, setProject] = useState({});
  const [owner, setOwner] = useState({});
  const [backings, setBackings] = useState();
  const [wosList, setWosList] = useState();

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const projectResponse = await projectService.getOne(params.id);
        setProject(projectResponse.project);
        setBackings(projectResponse.backings);
        setOwner(projectResponse.user);

        setWosList(
          projectResponse.backings.filter((b) => b.word_of_support !== null)
        );
      } catch (err) {
        navigate("/not-found");
        console.error(err);
      }
    };

    fetchProject();
  }, [params.id, navigate]);

  return (
    <div className="mx-auto">
      <div className="text-3xl font-bold m-5">{project.title}</div>
      <div className="grid grid-cols-3 gap-4 ">
        <div className="col-span-2">
          <div className="rounded-xl overflow-hidden h-128 aspect-[4/3] object-cover">
            <img
              src={project.cover_picture}
              className="w-full h-full m-auto object-cover"
            />
          </div>
          <div className="m-3 flex justify-start">
            <Avatar
              alt="User settings"
              img={
                owner.profile_picture ? owner.profile_picture : "/avatar.svg"
              }
              rounded
            >
              <div>
                <span className="flex justify-start tracking">
                  {`${owner.first_name} ${owner.last_name}`} is organizing this
                  fundraiser.
                </span>
                <span className="block font-normal text-gray-500 text-sm">
                  Created{" "}
                  {utils.calculateDayDifference(
                    Date.now(),
                    utils.parseDateFromRFC3339(project.start_date)
                  )}
                  d ago.
                </span>
              </div>
            </Avatar>
          </div>
          <div className="h-px bg-gray-300"></div>

          <p className="max-w-2xl tracking-tight mt-4">{project.description}</p>
          <div className="flex justify-center my-5">
            <Link to={`/${params.id}/donate`}>
              <div className="border text-xl flex py-3 max-w-lg rounded-lg border-gray-400 px-40 hover:bg-gray-100 hover:border-gray-900 duration-300">
                Donate
              </div>
            </Link>
          </div>

          <div className="h-px bg-gray-300"></div>

          <div className="my-5">
            <div className="text-xl font-semibold tracking-tight">
              Donors&apos; words of support ({wosList?.length})
            </div>

            {wosList?.map((b) => (
              <Support
                key={b.id}
                avatar={b.profile_picture ? b.profile_picture : "/avatar.svg"}
                first_name={b.first_name ? b.first_name : "Anonymous"}
                last_name={b.last_name ? b.last_name : ""}
                amount={b.amount}
                day_since={utils.getDaySince(b.created_at)}
                comment={b.word_of_support}
              />
            ))}
          </div>
        </div>
        <div className="col-span-1">
          <DonateBox
            id={params.id}
            currentAmount={project.current_amount}
            goalAmount={project.goal_amount}
            backings={backings}
          />
        </div>
      </div>
    </div>
  );
};

export default Project;
