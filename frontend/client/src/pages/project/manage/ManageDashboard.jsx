import { useState } from "react";
import utils from "../../../utils/utils";
import { Button, Tooltip } from "flowbite-react";
import { CiEdit } from "react-icons/ci";
import { FaRegCopy } from "react-icons/fa";

const ManageDashboard = () => {
  const [currentAmount] = useState(2000000);
  return (
    <div>
      <div className="text-3xl font-semibold">my son go to school</div>
      <div className="my-4 bg-gray-200 rounded-full h-1">
        <div
          className="bg-green-400 h-1 rounded-full"
          style={{
            width: `${utils.calculateProgress(2000, 5000)}%`,
          }}
        ></div>
      </div>
      <div className="flex justify-between">
        <div className="flex space-x-1">
          {currentAmount ? (
            <div>
              <span className="font-medium">
                ₫{utils.formatNumber(currentAmount)}
              </span>
              <span className="ml-1">raised of</span>
            </div>
          ) : (
            ""
          )}
          <div className="font-medium">₫{utils.formatNumber(5000000)} goal</div>
        </div>
      </div>

      <div className="flex space-x-2 mt-5">
        <Tooltip
          content="You can edit your fundraiser if it has not received any donation."
          style="light"
          placement="bottom"
        >
          <Button color="light" disabled={currentAmount > 0}>
            <CiEdit className="mr-2 h-5 w-5" />
            Edit
          </Button>
        </Tooltip>

        <Button
          color="light"
          onClick={() => navigator.clipboard.writeText("lmao xd")}
        >
          <FaRegCopy className="mr-2 h-5 w-5" />
          Copy fundraiser link
        </Button>
      </div>
    </div>
  );
};

export default ManageDashboard;
