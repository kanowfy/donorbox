import { Avatar } from "flowbite-react";
import PropTypes from "prop-types";
import utils from "../utils/utils";

const ProjectUpdate = ({
  profile_picture,
  first_name,
  last_name,
  content,
  created_at,
}) => {
  return (
    <div className="grid grid-cols-12">
      <div className="col-span-1">
        <Avatar
          img={profile_picture ? profile_picture : "/avatar.svg"}
          rounded
          size="sm"
        />
      </div>
      <div className="col-span-11">
        <div className="text-sm font-medium">
          {first_name} {last_name}
        </div>
        <div className="text-sm text-gray-600">
          {utils.getDaySince(created_at)}d
        </div>
        <div className="text-sm tracking-tight">{content}</div>
      </div>
    </div>
  );
};

ProjectUpdate.propTypes = {
  profile_picture: PropTypes.string,
  first_name: PropTypes.string,
  last_name: PropTypes.string,
  content: PropTypes.string,
  created_at: PropTypes.string,
};

export default ProjectUpdate;
