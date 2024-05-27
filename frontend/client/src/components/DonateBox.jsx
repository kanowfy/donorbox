import { Link } from "react-router-dom";
import utils from "../utils/utils";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";
import backingService from "../services/backing";
import { Avatar, Modal } from "flowbite-react";
import Donator from "./Donator";

const DonateBox = ({ id, currentAmount, goalAmount, backings }) => {
  const [recentBacking, setRecentBacking] = useState({});
  const [mostBacking, setMostBacking] = useState({});
  const [firstBacking, setFirstBacking] = useState({});
  const [backingCount, setBackingCount] = useState(0);
  const [viewDonations, setViewDonations] = useState(false);

  useEffect(() => {
    const fetchBackingStats = async (projectID) => {
      try {
        const response = await backingService.getProjectStats(projectID);
        console.log(response);
        setRecentBacking(response.recent_backing);
        setMostBacking(response.most_backing);
        setFirstBacking(response.first_backing);
        setBackingCount(response.backing_count);
      } catch (err) {
        console.error(err);
      }
    };

    fetchBackingStats(id);
  }, [id]);

  return (
    <div className="shadow-xl p-6 rounded-lg">
      <div>
        <Link to={`/${id}/donate`}>
          <div className="py-3 px-10 bg-gradient-to-b from-yellow-300 to-yellow-400 hover:from-yellow-400 hover:to-yellow-300 flex items-center justify-center font-semibold text-lg rounded-xl">
            Donate
          </div>
        </Link>
      </div>
      <div className="flex my-3 gap-2 items-end justify-end mx-2 tracking-tight">
        <span className="block text-xl font-medium">
          ₫{currentAmount?.toLocaleString()}
        </span>
        <span className="text-sm text-center block text-gray-700">
          of ₫{goalAmount?.toLocaleString()} raised
        </span>
      </div>
      <div className="mx-2 my-2 bg-gray-200 rounded-full h-1.5 dark:bg-gray-700">
        <div
          className="bg-green-500 h-1.5 rounded-full"
          style={{
            width: `${utils.calculateProgress(currentAmount, goalAmount)}%`,
          }}
        ></div>
      </div>
      <div className="flex justify-end font-sans text-sm antialiased font-normal text-gray-700 mx-2">
        {backingCount} donations
      </div>

      {backingCount ? (
        <div className="mt-7 mb-2 space-y-2">
          <div className="grid grid-cols-12">
            <div className="col-span-2">
              <Avatar
                alt="avatar"
                img={
                  recentBacking.profile_picture
                    ? recentBacking.profile_picture
                    : "/avatar.svg"
                }
                rounded
              />
            </div>
            <div className="col-span-10 flex flex-col">
              <div className="font-normal">
                {recentBacking.first_name
                  ? `${recentBacking.first_name} ${recentBacking.last_name}`
                  : "Anonymous"}
              </div>
              <div className="flex gap-2 my-1">
                <span className="text-sm block font-bold">
                  ₫{recentBacking?.amount?.toLocaleString()}{" "}
                </span>
                <span className="block">&#8226;</span>
                <span className="text-sm block text-gray-500">
                  {utils.getDaySince(recentBacking.created_at)}d
                </span>
                <span className="block">&#8226;</span>
                <span className="text-sm block text-gray-500">
                  Recent donation
                </span>
              </div>
            </div>
          </div>

          <hr />

          <div className="grid grid-cols-12">
            <div className="col-span-2">
              <Avatar
                alt="avatar"
                img={
                  mostBacking.profile_picture
                    ? mostBacking.profile_picture
                    : "/avatar.svg"
                }
                rounded
              />
            </div>
            <div className="col-span-10 flex flex-col">
              <div className="font-normal">
                {mostBacking.first_name
                  ? `${mostBacking.first_name} ${mostBacking.last_name}`
                  : "Anonymous"}
              </div>
              <div className="flex gap-2 my-1">
                <span className="text-sm block font-bold">
                  ₫{mostBacking?.amount?.toLocaleString()}{" "}
                </span>
                <span className="block">&#8226;</span>
                <span className="text-sm block text-gray-500">
                  {utils.getDaySince(mostBacking.created_at)}d
                </span>
                <span className="block">&#8226;</span>
                <span className="text-sm block text-gray-500">
                  Most donation
                </span>
              </div>
            </div>
          </div>

          <hr />

          <div className="grid grid-cols-12">
            <div className="col-span-2">
              <Avatar
                alt="avatar"
                img={
                  firstBacking.profile_picture
                    ? firstBacking.profile_picture
                    : "/avatar.svg"
                }
                rounded
              />
            </div>
            <div className="col-span-10 flex flex-col">
              <div className="font-normal">
                {firstBacking.first_name
                  ? `${firstBacking.first_name} ${firstBacking.last_name}`
                  : "Anonymous"}
              </div>
              <div className="flex gap-2 my-1">
                <span className="text-sm block font-bold">
                  ₫{firstBacking?.amount?.toLocaleString()}{" "}
                </span>
                <span className="block">&#8226;</span>
                <span className="text-sm block text-gray-500">
                  {utils.getDaySince(firstBacking.created_at)}d
                </span>
                <span className="block">&#8226;</span>
                <span className="text-sm block text-gray-500">
                  First donation
                </span>
              </div>
            </div>
          </div>

          <div className="w-full flex justify-center">
            <button
              className="hover:underline text-sm"
              onClick={() => setViewDonations(true)}
            >
              View all donations
            </button>
          </div>

          <Modal
            show={viewDonations}
            size="md"
            onClose={() => setViewDonations(false)}
            popup
          >
            <Modal.Header />
            <Modal.Body>
              <div className="mb-4 underline">List of donations</div>
              <div className="space-y-1">
                {backings?.map((b) => (
                  <Donator
                    key={b.id}
                    profile_picture={
                      b.profile_picture ? b.profile_picture : "/avatar.svg"
                    }
                    first_name={b.first_name ? b.first_name : "Anonymous"}
                    last_name={b.last_name ? b.last_name : ""}
                    amount={b.amount}
                    created_at={b.created_at}
                  />
                ))}
              </div>
            </Modal.Body>
          </Modal>
        </div>
      ) : (
        <div className="flex items-center justify-center space-x-1 mt-5">
          <div>Be the first to help!</div>
          <Link
            to={`/${id}/donate`}
            className="font-semibold text-gray-600 underline hover:text-black"
          >
            Donate now
          </Link>
        </div>
      )}
    </div>
  );
};

DonateBox.propTypes = {
  id: PropTypes.string,
  currentAmount: PropTypes.number,
  goalAmount: PropTypes.number,
  backings: PropTypes.array,
};

export default DonateBox;
