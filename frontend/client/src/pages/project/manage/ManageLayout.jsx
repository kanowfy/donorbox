import { Outlet, useParams, useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";
import ManageProjectSidebar from "../../../components/ManageProjectSidebar";
import projectService from "../../../services/project";

const ManageLayout = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [project, setProject] = useState();
  const [milestones, setMilestones] = useState();
  const [backings, setBackings] = useState();

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const response = await projectService.getOne(params.id);
        setProject(response.project);
        setMilestones(response.milestones);
        setBackings(response.backings);
      } catch (err) {
        navigate("/not-found");
        console.error(err);
      }
    };

    fetchProject();
  }, [params.id, navigate]);
  return (
    <section className="flex justify-center mt-14">
      <div className="grid grid-cols-10 gap-2 w-2/3">
        <div className="col-span-3 flex">
          <ManageProjectSidebar id={params.id} />
        </div>
        <div className="col-span-7 space-y-10">
          <Outlet context={{ project, milestones, backings }} />
        </div>
      </div>
    </section>
  );
};

export default ManageLayout;
