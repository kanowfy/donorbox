import { Avatar } from "flowbite-react";
import PropTypes from "prop-types";
import utils from "../utils/utils";

const ProjectUpdate = ({
  profile_picture,
  first_name,
  last_name,
  photo,
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
        {photo && (
          <div className="rounded-xl overflow-hidden h-40 aspect-[4/3] object-cover my-4">
            <img src={photo} className="w-full h-full m-auto object-cover" />
          </div>
        )}
        <div className="text-sm tracking-tight">{content}</div>
      </div>
    </div>
  );
};

ProjectUpdate.propTypes = {
  profile_picture: PropTypes.string,
  first_name: PropTypes.string,
  last_name: PropTypes.string,
  photo: PropTypes.string,
  content: PropTypes.string,
  created_at: PropTypes.string,
};

export default ProjectUpdate;
