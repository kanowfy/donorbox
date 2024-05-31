import utils from "../../../utils/utils";
import { Button, Tooltip } from "flowbite-react";
import { CiEdit } from "react-icons/ci";
import { FaRegCopy } from "react-icons/fa";
import { MdPayment } from "react-icons/md";
import { Link, useOutletContext } from "react-router-dom";

const ManageDashboard = () => {
  const { project } = useOutletContext();
  return (
    <div>
      <div className="rounded-xl overflow-hidden h-96 aspect-[4/3] object-cover">
        <img
          src={project?.cover_picture}
          className="w-full h-full m-auto object-cover"
        />
      </div>
      <div className="text-3xl font-semibold mt-2">{project?.title}</div>
      <div className="my-4 bg-gray-200 rounded-full h-1">
        <div
          className="bg-green-400 h-1 rounded-full"
          style={{
            width: `${utils.calculateProgress(
              project?.current_amount,
              project?.goal_amount
            )}%`,
          }}
        ></div>
      </div>
      <div className="flex justify-between">
        <div className="flex space-x-1">
          {project?.current_amount > 0 && (
            <div>
              <span className="font-medium">
                ₫
                {project?.current_amount &&
                  utils.formatNumber(project?.current_amount)}
              </span>
              <span className="ml-1">raised of</span>
            </div>
          )}
          <div className="font-medium">
            ₫{project?.goal_amount && utils.formatNumber(project?.goal_amount)}{" "}
            goal
          </div>
        </div>
      </div>

      <div className="flex space-x-2 mt-5">
        <Tooltip
          content="You can edit your fundraiser if it has not received any donation."
          style="light"
          placement="bottom"
        >
          <Button color="light" disabled={project?.current_amount > 0}>
            <CiEdit className="mr-2 h-5 w-5" />
            Edit
          </Button>
        </Tooltip>

        <Tooltip
          content="Copy unique link to fundraiser"
          style="light"
          placement="bottom"
        >
          <Button
            color="light"
            onClick={() =>
              navigator.clipboard.writeText(
                `localhost:5173/fundraiser/${project?.id}`
              )
            }
          >
            <FaRegCopy className="mr-2 h-5 w-5" />
            Copy link
          </Button>
        </Tooltip>

        <Tooltip
          content="Setup transfer before due date so you can receive funds when your fundraiser succeeds."
          style="light"
          placement="bottom"
        >
          <Link to={`/manage/${project?.id}/transfer`}>
            <Button color="light" disabled={project?.card_id}>
              <MdPayment className="mr-2 h-5 w-5" />
              Setup Transfer
            </Button>
          </Link>
        </Tooltip>
      </div>
    </div>
  );
};

export default ManageDashboard;
