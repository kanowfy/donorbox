import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import utils from "../utils/utils";

const ProjectCard = (props) => {
  return (
    <Link to="#">
      <div className="relative flex flex-col mt-6 text-gray-700 bg-green-50 shadow-md bg-clip-border rounded-xl w-72">
        <div className="relative h-40 mx-4 -mt-6 overflow-hidden text-white shadow-lg bg-clip-border rounded-xl bg-blue-gray-500 shadow-blue-gray-500/40">
          <img src={props.cover} alt="card-image" />
        </div>
        <div className="px-4 pt-3">
          <div className="h-12 mb-5">
            <h5 className="block mb-2 font-sans text-base tracking-tight antialiased font-medium leading-tight text-blue-gray-900 line-clamp-2 text-ellipsis">
              {props.title}
            </h5>
          </div>
          <div className="flex justify-between">
            <p className="block font-sans text-sm antialiased font-normal">
              {props.numBackings} backings
            </p>
            <p className="block font-sans text-sm antialiased font-normal">
              ₫{props.currentAmount.toLocaleString()} donated
            </p>
          </div>
        </div>
        <div className="mx-2 my-2 bg-gray-200 rounded-full h-1.5 dark:bg-gray-700">
          <div
            className="bg-green-500 h-1.5 rounded-full"
            style={{
              width: `${utils.calculateProgress(
                props.currentAmount,
                props.goalAmount
              )}%`,
            }}
          ></div>
        </div>
      </div>
    </Link>
  );
};

ProjectCard.propTypes = {
  title: PropTypes.string.isRequired,
  cover: PropTypes.string.isRequired,
  currentAmount: PropTypes.number,
  goalAmount: PropTypes.number,
  numBackings: PropTypes.number,
};

export default ProjectCard;
