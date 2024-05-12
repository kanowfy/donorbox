import { Link } from "react-router-dom";
import utils from "../utils/utils";

const DonateBox = () => {
  return (
    <div className="shadow-xl p-6 rounded-lg">
      <div>
        <Link to="#">
          <div className="py-3 px-10 bg-gradient-to-b from-yellow-300 to-yellow-400 hover:from-yellow-400 hover:to-yellow-300 flex items-center justify-center font-semibold text-lg rounded-xl">
            Donate now
          </div>
        </Link>
      </div>
      <div className="flex my-3 text-gray-900 gap-2 justify-items-end mx-2 align-bottom">
        <span className="block text-xl">₫2,500,000</span>
        <span className="text-sm text-center block ">of ₫5,000,000 raised</span>
      </div>
      <div className="mx-2 my-2 bg-gray-200 rounded-full h-1.5 dark:bg-gray-700">
        <div
          className="bg-green-500 h-1.5 rounded-full"
          style={{
            width: `${utils.calculateProgress(2500, 5000)}%`,
          }}
        ></div>
      </div>
      <div className="block font-sans text-sm antialiased font-normal mx-2">
        3100 backings
      </div>
    </div>
  );
};

export default DonateBox;
