import { Link } from "react-router-dom";
import utils from "../utils/utils";
import PropTypes from "prop-types";
import { useEffect, useState } from "react";
import backingService from "../services/backing";
import { Avatar, Modal } from "flowbite-react";
import { PiShareFatThin } from "react-icons/pi";
import { FaFacebook } from "react-icons/fa";
import { FaXTwitter } from "react-icons/fa6";
import { IoIosMail } from "react-icons/io";
import { IoMdCopy } from "react-icons/io";
import Donor from "./Donor";
import { SERVE_URL } from "../constants";

const DonateBox = ({ id, totalFund, fundGoal, backings }) => {
  const [recentBacking, setRecentBacking] = useState({});
  const [mostBacking, setMostBacking] = useState({});
  const [firstBacking, setFirstBacking] = useState({});
  const [backingCount, setBackingCount] = useState(0);
  const [viewDonations, setViewDonations] = useState(false);
  const [isOpenShare, setIsOpenShare] = useState(false);

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
    <div className="w-full shadow-xl p-6 rounded-lg">
      <div className="space-y-2">
        <Link to={`/fundraiser/${id}/donate`}>
          <div className="py-3 px-10 bg-gradient-to-b from-yellow-300 to-yellow-400 hover:from-yellow-400 hover:to-yellow-300 flex items-center justify-center font-semibold text-lg rounded-xl">
            Donate
          </div>
        </Link>
        <div
          className="cursor-pointer hover:bg-gray-200 py-3 px-10 border-gray-900 border flex items-center justify-center font-semibold text-lg rounded-xl"
          onClick={() => setIsOpenShare(true)}
        >
          Share
        </div>
      </div>
      <div className="flex flex-col my-5 space-y-1 justify-center items-center">
        <div className="block text-4xl font-bold text-yellow-400">
          ₫{totalFund?.toLocaleString()}
        </div>
        <div className="text-gray-700">
          raised of{" "}
          <span className="text-yellow-500">₫{fundGoal?.toLocaleString()}</span>{" "}
          target
        </div>
        <div className="text-gray-700">
          from <span className="text-teal-500">{backingCount}</span> donations
        </div>
      </div>

      {backingCount ? (
        <div className="mt-7 mb-2 space-y-2">
          <div className="grid grid-cols-12">
            <div className="col-span-2">
              <Avatar
                alt="avatar"
                img={
                  recentBacking.backer.profile_picture
                    ? recentBacking.backer.profile_picture
                    : "/avatar.svg"
                }
                rounded
              />
            </div>
            <div className="col-span-10 flex flex-col">
              <div className="font-normal">
                {recentBacking.backer.first_name
                  ? `${recentBacking.backer.first_name} ${recentBacking.backer.last_name}`
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
                  mostBacking.backer.profile_picture
                    ? mostBacking.backer.profile_picture
                    : "/avatar.svg"
                }
                rounded
              />
            </div>
            <div className="col-span-10 flex flex-col">
              <div className="font-normal">
                {mostBacking.backer.first_name
                  ? `${mostBacking.backer.first_name} ${mostBacking.backer.last_name}`
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
                  firstBacking.backer.profile_picture
                    ? firstBacking.backer.profile_picture
                    : "/avatar.svg"
                }
                rounded
              />
            </div>
            <div className="col-span-10 flex flex-col">
              <div className="font-normal">
                {firstBacking.backer.first_name
                  ? `${firstBacking.backer.first_name} ${firstBacking.backer.last_name}`
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
                  <Donor
                    key={b.id}
                    profile_picture={
                      b.backer.profile_picture
                        ? b.backer.profile_picture
                        : "/avatar.svg"
                    }
                    first_name={
                      b.backer.first_name ? b.backer.first_name : "Anonymous"
                    }
                    last_name={b.backer.last_name ? b.backer.last_name : ""}
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
            to={`/fundraiser/${id}/donate`}
            className="font-semibold text-gray-600 underline hover:text-black"
          >
            Donate now
          </Link>
        </div>
      )}
      <Modal show={isOpenShare} onClose={() => setIsOpenShare(false)} size="xl">
        <Modal.Header>
          <div className="flex">
            Share this Campaign
            <PiShareFatThin className="w-7 h-7 ml-2" />
          </div>
        </Modal.Header>
        <Modal.Body>
          <div className="grid grid-cols-4 items-center justify-center leading-tight text-gray-700 px-10">
            <a
              className="flex flex-col items-center space-y-1"
              href={`mailto:?body=${encodeURI(
                `${SERVE_URL}/fundraiser/${id}`
              )}`}
              target="_blank"
            >
              <IoIosMail className="text-red-500 w-16 h-16" />
              <div>Send email</div>
            </a>
            <a
              className="flex flex-col items-center space-y-1"
              href={`https://www.facebook.com/sharer/sharer.php?u=${encodeURI(
                `${SERVE_URL}/fundraiser/${id}`
              )}`}
              target="_blank"
            >
              <FaFacebook className="text-blue-500 w-16 h-16" />
              <div>Facebook</div>
            </a>
            <a
              className="flex flex-col items-center space-y-1"
              href={`https://x.com/intent/tweet?url=${encodeURI(
                `${SERVE_URL}/fundraiser/${id}`
              )}`}
              target="_blank"
            >
              <FaXTwitter className="text-zinc-800 w-16 h-16" />
              <div>X</div>
            </a>
            <div
              className="flex flex-col items-center space-y-1 cursor-pointer"
              onClick={() => {
                navigator.clipboard.writeText(`${SERVE_URL}/fundraiser/${id}`);
              }}
            >
              <IoMdCopy className=" w-16 h-16" />
              <div>Copy link</div>
            </div>
          </div>
        </Modal.Body>
      </Modal>
    </div>
  );
};

DonateBox.propTypes = {
  id: PropTypes.string,
  totalFund: PropTypes.number,
  fundGoal: PropTypes.number,
  backings: PropTypes.array,
};

export default DonateBox;
