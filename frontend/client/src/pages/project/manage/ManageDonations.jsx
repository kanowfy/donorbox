import Donator from "../../../components/Donator";

const ManageDonations = () => {
  return (
    <div>
      <div className="text-3xl font-semibold">Donations</div>
      <div className="text-gray-600 text-sm mt-1">
        See all donations in your fundraiser
      </div>
      <div className="my-10 grid grid-cols-2 gap-5">
        <Donator
          first_name="John"
          last_name="Doe"
          created_at="2024-05-21T07:30:37+00:00"
          amount={90000}
        />
        <Donator
          first_name="John"
          last_name="Doe"
          created_at="2024-05-21T07:30:37+00:00"
          amount={90000}
        />
        <Donator
          first_name="John"
          last_name="Doe"
          created_at="2024-05-21T07:30:37+00:00"
          amount={90000}
        />
      </div>
    </div>
  );
};

export default ManageDonations;
