import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import utils from "../utils/utils";

const ProjectCard = ({
  id,
  cover,
  title,
  totalFund,
  fundGoal,
  numBackings,
}) => {
  return (
    <Link to={`/fundraiser/${id}`}>
      <div className="relative flex flex-col mt-6 text-gray-700 bg-green-50 shadow-md bg-clip-border rounded-xl w-72">
        <div className="relative h-40 mx-4 -mt-6 overflow-hidden text-white shadow-lg bg-clip-border rounded-xl bg-blue-gray-500 shadow-blue-gray-500/40">
          <img src={cover} alt="card-image" />
        </div>
        <div className="px-4 pt-3">
          <div className="h-12 mb-5">
            <h5 className="block mb-2 text-base tracking-tight antialiased font-medium leading-tight text-blue-gray-900 line-clamp-2 text-ellipsis">
              {title}
            </h5>
          </div>
          <div className="flex justify-between">
            <p className="block text-sm antialiased font-medium text-gray-600">
              {numBackings} donations
            </p>
            <p className="block text-sm antialiased font-medium text-gray-600">
              ₫{totalFund.toLocaleString()} raised
            </p>
          </div>
        </div>
        <div className="mx-2 my-2 bg-gray-200 rounded-full h-1.5">
          <div
            className="bg-blue-500 h-1.5 rounded-full"
            style={{
              width: `${utils.calculateProgress(totalFund, fundGoal)}%`,
            }}
          ></div>
        </div>
      </div>
    </Link>
  );
};

ProjectCard.propTypes = {
  id: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
  cover: PropTypes.string.isRequired,
  totalFund: PropTypes.number,
  fundGoal: PropTypes.number,
  numBackings: PropTypes.number,
};

export default ProjectCard;
