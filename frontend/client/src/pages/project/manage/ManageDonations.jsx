import Donator from "../../../components/Donator";
import { useOutletContext } from "react-router-dom";

const ManageDonations = () => {
  const { backings } = useOutletContext();
  return (
    <div>
      <div className="text-3xl font-semibold">Donations</div>
      <div className="text-gray-600 text-sm mt-1">
        See all donations in your fundraiser
      </div>
      <div className="my-10 grid grid-cols-2 gap-5">
        {backings?.map((b) => (
          <Donator
            key={b.id}
            profile_picture={b?.profile_picture}
            first_name={b.first_name ? b.first_name : "Anonymous"}
            last_name={b.last_name ? b.last_name : ""}
            amount={b.amount}
            created_at={b.created_at}
          />
        ))}
      </div>
    </div>
  );
};

export default ManageDonations;
