import Donor from "../../../components/Donor";
import { useOutletContext } from "react-router-dom";

const ManageDonations = () => {
  const { backings } = useOutletContext();
  return (
    <div>
      <div className="text-3xl font-semibold">Donations</div>
      <div className="text-gray-600 text-sm mt-1">
        See all donations in your fundraiser
      </div>
      {
        backings ?
      (<div className="my-10 grid grid-cols-2 gap-5">
        {backings.map((b) => (
          <Donor
            key={b.id}
            profile_picture={b?.backer.profile_picture}
            first_name={b.backer.first_name ? b.backer.first_name : "Anonymous"}
            last_name={b.backer.last_name ? b.backer.last_name : ""}
            amount={b.amount}
            created_at={b.created_at}
          />
        ))}
      </div>) : (
        <div className="my-5 text-2xl text-gray-800">This project has not received any funds yet.</div>
      )
      }
    </div>
  );
};

export default ManageDonations;
