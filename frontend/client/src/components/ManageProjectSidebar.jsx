import { Sidebar } from "flowbite-react";
import { IoMdHome } from "react-icons/io";
import { FaClover } from "react-icons/fa6";
import { MdChat } from "react-icons/md";
import { GrDocumentUser } from "react-icons/gr";
import PropTypes from "prop-types";
import { Link } from "react-router-dom";

const ManageProjectSidebar = ({ id }) => {
  return (
    <Sidebar>
      <Sidebar.Items>
        <Sidebar.ItemGroup>
          <Sidebar.Item as={Link} to={`/manage/${id}`} icon={IoMdHome}>
            <span className="font-medium">Dashboard</span>
          </Sidebar.Item>
          <Sidebar.Item
            as={Link}
            to={`/manage/${id}/donations`}
            icon={FaClover}
          >
            <span className="font-medium">Donations</span>
          </Sidebar.Item>
          <Sidebar.Item as={Link} to={`/manage/${id}/proofs`} icon={GrDocumentUser}>
            <span className="font-medium">Proofs</span>
          </Sidebar.Item>
        </Sidebar.ItemGroup>
      </Sidebar.Items>
    </Sidebar>
  );
};

ManageProjectSidebar.propTypes = {
  id: PropTypes.string,
};

export default ManageProjectSidebar;
