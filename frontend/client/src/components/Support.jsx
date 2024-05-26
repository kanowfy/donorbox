import { Avatar } from "flowbite-react";
import PropTypes from "prop-types";

const Support = ({ avatar, amount, day_since, comment }) => {
  return (
    <div className="grid grid-cols-12 my-7">
      <div className="col-span-1">
        <Avatar alt="avatar" img={avatar} rounded />
      </div>
      <div className="col-span-11 flex flex-col">
        <div className="text-base font-semibold">Jack Sparrow</div>
        <div className="flex gap-2 my-1">
          <span className="text-sm block">₫{amount.toLocaleString()} </span>
          <span className="block">&#8226;</span>
          <span className="text-sm block text-gray-500">{day_since}d</span>
        </div>
        <div className="max-w-2xl text-base leading-tight tracking-tight">
          {comment}
        </div>
      </div>
    </div>
  );
};

Support.propTypes = {
  avatar: PropTypes.string,
  amount: PropTypes.number,
  day_since: PropTypes.number,
  comment: PropTypes.string,
};

export default Support;
