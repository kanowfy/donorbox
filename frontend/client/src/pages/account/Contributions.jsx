import { useEffect, useState } from "react";
import { useAuthContext } from "../../context/AuthContext";
import { Button, Card } from "flowbite-react";
import { Link } from "react-router-dom";
import utils from "../../utils/utils";
import backingService from "../../services/backing";

const Contributions = () => {
  const { token } = useAuthContext();
  const [backings, setBackings] = useState();

  useEffect(() => {
    const fetchBackings = async (token) => {
      try {
        const response = await backingService.getBackingsForUser(token);
        setBackings(response.backings);
        console.log(response.backings);
      } catch (err) {
        console.error(err);
      }
    };

    fetchBackings(token);
  });

  return (
    <section className="py-10 flex flex-col items-center">
      <div className="flex justify-between w-2/5">
        <div className="text-3xl font-semibold my-2">Donation History</div>
      </div>
      <div className="w-2/5">
        {backings ? (
          <div className="mt-10 space-y-3">
            {backings.map((b) => (
              <div
                className="rounded-lg border py-5 px-4 grid grid-cols-10 gap-4"
                key={b.id}
              >
                <div className="overflow-hidden col-span-3">
                  <Link to={`/fundraiser/${b.project_id}`}>
                    <img
                      className="rounded-lg aspect-[4/3] h-32"
                      src={b.cover_picture}
                    />
                  </Link>
                </div>
                <div className="col-span-7 flex flex-col justify-between">
                  <div className="space-y-4">
                    <div className="flex justify-between">
                      <div className="font-medium">{b.title}</div>
                      <div className="font-semibold mr-4">
                        ₫{b.amount.toLocaleString()}
                      </div>
                    </div>
                    <div className="text-sm text-gray-700">
                      on{" "}
                      {utils.formatDate(
                        new Date(utils.parseDateFromRFC3339(b.created_at))
                      )}
                    </div>
                  </div>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="flex flex-col items-center space-y-3">
            <div className="text-xl mt-20">You have not made any donation.</div>
            <Link
              to="/"
              className="hover:text-black text-gray-800 hover:underline hover:bg-gray-100 px-3 py-2 border rounded-md border-gray-700"
            >
              Back to Homepage
            </Link>
          </div>
        )}
      </div>
    </section>
  );
};

export default Contributions;
