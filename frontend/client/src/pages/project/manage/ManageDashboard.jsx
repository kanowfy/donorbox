import utils from "../../../utils/utils";
import { Button, Timeline, Tooltip } from "flowbite-react";
import { CiEdit } from "react-icons/ci";
import { FiEye } from "react-icons/fi";
import { FaRegCopy } from "react-icons/fa";
import { HiCalendar } from "react-icons/hi";
import { useOutletContext, useNavigate } from "react-router-dom";
import { SERVE_URL } from "../../../constants";

const ManageDashboard = () => {
  const navigate = useNavigate();
  const { project, milestones } = useOutletContext();
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
              project?.total_fund,
              project?.fund_goal
            )}%`,
          }}
        ></div>
      </div>
      <div className="flex justify-between">
        <div className="flex space-x-1 text-xl">
          {project?.total_fund > 0 && (
            <div>
              <span className="font-medium">
                ₫
                {project?.total_fund && utils.formatNumber(project?.total_fund)}
              </span>
              <span className="ml-1">raised of</span>
            </div>
          )}
          <div className="font-medium">
            ₫{project?.fund_goal && utils.formatNumber(project?.fund_goal)} goal
          </div>
        </div>
      </div>

      <div className="mt-10">
        <Timeline horizontal>
          {milestones
            ?.sort((a, b) => a.id - b.id)
            .map((m) => (
              <Timeline.Item key={m.id}>
                <Timeline.Point icon={HiCalendar} />
                <Timeline.Content>
                  <Timeline.Title>{m.title}</Timeline.Title>
                  <div className=" text-red-500 tracking-tight font-semibold">
                    ₫{utils.formatNumber(m.current_fund)} / ₫
                    {utils.formatNumber(m.fund_goal)}
                  </div>
                  <Timeline.Body>{m?.description}</Timeline.Body>
                </Timeline.Content>
              </Timeline.Item>
            ))}
        </Timeline>
      </div>
      <div className="flex space-x-2 mt-5">
        <Tooltip
          content="Go to milestone page."
          style="light"
          placement="bottom"
        >
          <Button
            color="light"
            onClick={() => {
              navigate(`/fundraiser/${project?.id}`);
            }}
          >
            <FiEye className="mr-2 h-5 w-5" />
            View
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
                `${SERVE_URL}/fundraiser/${project?.id}`
              )
            }
          >
            <FaRegCopy className="mr-2 h-5 w-5" />
            Copy link
          </Button>
        </Tooltip>
      </div>
    </div>
  );
};

export default ManageDashboard;
