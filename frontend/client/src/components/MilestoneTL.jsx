import { Timeline, Tooltip, Badge } from "flowbite-react";
import { HiCalendar } from "react-icons/hi";
import { IoEyeOutline } from "react-icons/io5";
import { MdDone } from "react-icons/md";
import utils from "../utils/utils";

const MilestoneStatusMap = {
  fund_released: {
    color: "info",
    text: "fund_released"
  },
  completed: {
    color: "success",
    text: "completed"
  },
  refuted: {
    color: "failure",
    text: "refuted"
  },
  pending: {
    color: "gray",
    text: "ongoing"
  }
}

const MilestoneTL = ({ milestones, setIsOpenMilestone, setMilestoneReview }) => {
  return (
    <Timeline>
      {milestones
        ?.sort((a, b) => a.id - b.id)
        .map((m) => (
          <Timeline.Item key={m.id}>
            <Timeline.Point icon={HiCalendar} />
            <Timeline.Content>
              <div className="flex justify-between">
              <Timeline.Title
                className={m.status === "completed" ? "line-through text-gray-500" : ""}
              >
                {m.title}
              </Timeline.Title>
              <Badge color={MilestoneStatusMap[m.status].color}>{MilestoneStatusMap[m.status].text}</Badge>

              </div>
              <div
                className={`${
                  m.status === "completed" ? "line-through text-gray-500" : "text-red-600"
                } tracking-tight font-semibold`}
              >
                ₫{utils.formatNumber(m.current_fund)} / ₫
                {utils.formatNumber(m.fund_goal)}
              </div>
              <Timeline.Body>{m?.description}</Timeline.Body>
              {["fund_released", "completed", "refuted"].includes(m.status) && (
                <Tooltip content="Click to view milestone progress">
                  <div onClick={() => {
                    setMilestoneReview(m);
                    setIsOpenMilestone(true);
                  }} className="flex cursor-pointer text-green-500 hover:bg-green-100 px-3 py-1 border rounded-full border-green-300">
                    <div className="text-sm text-cyan-800">View Details</div>
                    {/*<MdDone className="ml-1 mt-1 w-5 h-5" />*/}
                  </div>
                </Tooltip>
              )}
            </Timeline.Content>
          </Timeline.Item>
        ))}
    </Timeline>
  );
};

export default MilestoneTL;
