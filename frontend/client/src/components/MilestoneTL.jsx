import { Timeline, Tooltip } from "flowbite-react";
import { HiCalendar } from "react-icons/hi";
import { IoEyeOutline } from "react-icons/io5";
import { MdDone } from "react-icons/md";
import utils from "../utils/utils";

const MilestoneTL = ({ milestones, setIsOpenMilestone, setMilestoneReview }) => {
  return (
    <Timeline>
      {milestones
        ?.sort((a, b) => a.id - b.id)
        .map((m) => (
          <Timeline.Item key={m.id}>
            <Timeline.Point icon={HiCalendar} />
            <Timeline.Content>
              <Timeline.Title
                className={m.completed ? "line-through text-gray-500" : ""}
              >
                {m.title}
              </Timeline.Title>
              <div
                className={`${
                  m.completed ? "line-through text-gray-500" : "text-red-500"
                } tracking-tight font-semibold`}
              >
                ₫{utils.formatNumber(m.current_fund)} / ₫
                {utils.formatNumber(m.fund_goal)}
              </div>
              <Timeline.Body>{m?.description}</Timeline.Body>
              {m.completed && (
                <Tooltip content="Click to view completion details">
                  <div onClick={() => {
                    setMilestoneReview(m);
                    setIsOpenMilestone(true);
                  }} className="flex cursor-pointer text-green-500 hover:bg-green-100 pl-3 pr-2 py-1 border rounded-full border-green-300">
                    <div>Resolved</div>
                    <MdDone className="ml-1 mt-1 w-5 h-5" />
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
