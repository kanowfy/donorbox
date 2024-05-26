import { Avatar } from "flowbite-react";
import utils from "../utils/utils";
import PropTypes from "prop-types";

const Donator = ({
  profile_picture,
  first_name,
  last_name,
  amount,
  created_at,
}) => {
  return (
    <div className="grid grid-cols-10 border border-gray-800 rounded-lg py-2 w-full">
      <div className="col-span-2">
        <Avatar
          alt="avatar"
          img={profile_picture ? profile_picture : "/avatar.svg"}
          rounded
        />
      </div>
      <div className="col-span-8 flex flex-col">
        <div className="font-normal">{`${first_name} ${last_name}`}</div>
        <div className="flex gap-2 my-1">
          <span className="text-sm block font-bold">
            ₫{amount.toLocaleString()}{" "}
          </span>
          <span className="block">&#8226;</span>
          <span className="text-sm block text-gray-500">
            {utils.getDaySince(created_at)}d
          </span>
        </div>
      </div>
    </div>
  );
};

Donator.propTypes = {
  profile_picture: PropTypes.string,
  first_name: PropTypes.string,
  last_name: PropTypes.string,
  amount: PropTypes.number,
  created_at: PropTypes.string,
};

export default Donator;
