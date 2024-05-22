import { Outlet } from "react-router-dom";
import ManageProjectSidebar from "../../../components/ManageProjectSidebar";

const ManageLayout = () => {
  return (
    <section className="flex justify-center mt-14">
      <div className="grid grid-cols-10 gap-2 w-2/3">
        <div className="col-span-3">
          <ManageProjectSidebar id="69" />
        </div>
        <div className="col-span-7 space-y-10">
          <Outlet />
        </div>
      </div>
    </section>
  );
};

export default ManageLayout;
