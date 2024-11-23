import { Timeline } from "flowbite-react";
import { HiCalendar } from "react-icons/hi";
import utils from "../utils/utils";

const MilestoneTL = ({milestones}) => {
    return (
        <Timeline>
            {milestones?.sort((a, b) => a.id -b.id).map((m) => (
            <Timeline.Item key={m.id}>
                <Timeline.Point icon={HiCalendar}/>
                <Timeline.Content>
                    <Timeline.Title>{m.title}</Timeline.Title>
                    <div className=" text-red-500 tracking-tight font-semibold">
                        ₫{utils.formatNumber(m.current_fund)} / ₫{utils.formatNumber(m.fund_goal)}
                    </div>
                    <Timeline.Body>
                        {m?.description}
                    </Timeline.Body>
                </Timeline.Content>
            </Timeline.Item>
                
            ))}
        </Timeline>
    )
}

export default MilestoneTL;