import { Button } from "flowbite-react";
import PropTypes from "prop-types";
import { Link } from "react-router-dom";

const ProjectListCard = ({ id, img, title, date }) => {
  return (
    <div className="rounded-lg border py-5 px-4 grid grid-cols-10 gap-4">
      <div className="overflow-hidden col-span-3">
        <Link to={`/manage/${id}`}>
          <img className="rounded-lg aspect-[4/3] h-32" src={img} />
        </Link>
      </div>
      <div className="col-span-7 flex flex-col justify-between">
        <div className="space-y-4">
          <div className="font-medium">{title}</div>
          <div className="text-sm text-gray-600">
            Fundraiser created on {date}
          </div>
        </div>
        <div className="flex space-x-2">
          <Link to={`/manage/${id}`}>
            <Button color="light">Manage</Button>
          </Link>
          <Link to={`/fundraiser/${id}`}>
            <Button color="light">View</Button>
          </Link>
        </div>
      </div>
    </div>
  );
};

ProjectListCard.propTypes = {
  id: PropTypes.string,
  img: PropTypes.string,
  title: PropTypes.string,
  date: PropTypes.string,
};

export default ProjectListCard;
