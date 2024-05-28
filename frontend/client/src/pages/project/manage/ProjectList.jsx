import { Button } from "flowbite-react";
import { useState, useEffect } from "react";
import ProjectListCard from "../../../components/ProjectListCard";
import utils from "../../../utils/utils";
import { useAuthContext } from "../../../context/AuthContext";
import projectService from "../../../services/project";
import { useNavigate, Link } from "react-router-dom";

const ProjectList = () => {
  const navigate = useNavigate();
  const [projects, setProjects] = useState();
  const { token } = useAuthContext() || {};

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const response = await projectService.getForUser(token);
        setProjects(response.projects);
      } catch (err) {
        console.error(err);
      }
    };

    if (token) {
      fetchProjects();
    } else {
      navigate("/login");
    }
  }, [token, navigate]);

  return (
    <section className="py-10 flex flex-col items-center">
      <div className="flex justify-between w-2/5">
        <div className="text-3xl font-semibold my-2">Your Fundraisers</div>
        <Link to="/start-fundraiser">
          <Button color="green" className="h-fit" size="lg">
            Start a Fundraiser
          </Button>
        </Link>
      </div>
      <div className="w-2/5">
        {projects ? (
          <div className="mt-20 space-y-3">
            {projects?.map((p) => (
              <ProjectListCard
                key={p.id}
                id={p.id}
                title={p.title}
                img={p.cover_picture}
                date={utils.formatDate(
                  new Date(utils.parseDateFromRFC3339(p.start_date))
                )}
              />
            ))}
          </div>
        ) : (
          <div className="flex justify-center text-lg mt-28">
            You have not started any fundraiser.
          </div>
        )}
      </div>
    </section>
  );
};

export default ProjectList;
