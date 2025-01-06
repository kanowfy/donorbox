import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import utils from "../utils/utils";
import {Badge} from "flowbite-react";

const ProjectCard = ({
  id,
  cover,
  title,
  totalFund,
  fundGoal,
  category,
  numBackings,
}) => {
  return (
    <Link to={`/fundraiser/${id}`}>
      <div className="relative flex flex-col mt-6 text-gray-700 bg-gray-100 shadow-lg border border-gray-300 bg-clip-border rounded-xl max-w-96">
        <div className="relative h-48 mx-4 mt-6 overflow-hidden text-white shadow-lg bg-clip-border rounded-xl bg-blue-gray-500 shadow-blue-gray-500/40">
          <img src={cover} alt="card-image" />
        </div>
        <div className="px-4 pt-3">
          <div className="h-12 mb-8">
            <div className="block mb-2 text-lg tracking-tight antialiased font-medium leading-tight text-blue-gray-900 line-clamp-2 text-ellipsis">
              {title}
            </div>
          {category && <Badge size="xs" className="w-fit">{category}</Badge>}
          </div>
          <div className="flex justify-between">
            <p className="block text-sm antialiased font-semibold text-cyan-700">
              {numBackings} donations
            </p>
            <p className="block text-sm antialiased font-semibold text-yellow-500">
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
  id: PropTypes.number.isRequired,
  title: PropTypes.string.isRequired,
  cover: PropTypes.string.isRequired,
  totalFund: PropTypes.number,
  fundGoal: PropTypes.number,
  numBackings: PropTypes.number,
};

export default ProjectCard;
