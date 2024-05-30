import {
  Table,
  TableHead,
  TableHeaderCell,
  TableBody,
  TableRow,
  TableCell,
  Card,
  Button,
  Dialog,
  DialogPanel,
  Badge,
} from "@tremor/react";
import PropTypes from "prop-types";
import { useState } from "react";
import { CategoryIndexMap } from "../../constants";
import { IoOpenOutline } from "react-icons/io5";
import utils from "../../utils/utils";
import fundService from "../../services/fund";
import { useNavigate } from "react-router-dom";

const PendingPayoutTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [review, setReview] = useState();
  const navigate = useNavigate();

  const handlePayout = async () => {
    try {
      console.log(token);
      await fundService.payout(token, review.id);
      setIsOpenReview(false);
      setIsSuccessful(true);
      setTimeout(() => {
        setIsSuccessful(false);
        navigate(0);
      }, 1500);
    } catch (err) {
      setIsFailed(true);
      console.error(err);
    }
  };

  return (
    <>
      <Card className="my-5 shadow-lg shadow-gray-500">
        <h3 className="text-tremor-content-strong dark:text-dark-tremor-content-strong font-semibold">
          List of projects eligible for payout
        </h3>
        <Table className="mt-5">
          <TableHead>
            <TableRow>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>Goal Amount</TableHeaderCell>
              <TableHeaderCell>Donated Amount</TableHeaderCell>
              <TableHeaderCell># Donations</TableHeaderCell>
              <TableHeaderCell>End date</TableHeaderCell>
              <TableHeaderCell>Transfer Setup</TableHeaderCell>
              <TableHeaderCell></TableHeaderCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((item) => (
              <TableRow key={item.id}>
                <TableCell className="font-medium text-black">
                  {item.title}
                </TableCell>
                <TableCell>{item.goal_amount.toLocaleString()}</TableCell>
                <TableCell>{item.current_amount.toLocaleString()}</TableCell>
                <TableCell>
                  <Badge>{item.backing_count}</Badge>
                </TableCell>
                <TableCell>
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(item.end_date))
                  )}
                </TableCell>
                <TableCell>
                  {item?.card_id ? (
                    <>
                      <Badge color="green">Yes</Badge>
                    </>
                  ) : (
                    <>
                      <Badge color="red">No</Badge>
                    </>
                  )}
                </TableCell>
                <TableCell>
                  <Button
                    variant="secondary"
                    onClick={() => {
                      setIsOpenReview(true);
                      setReview(item);
                    }}
                  >
                    Review
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Card>
      <Dialog
        open={isOpenReview}
        onClose={(val) => setIsOpenReview(val)}
        static={true}
      >
        <DialogPanel className="w-full">
          <div className="space-y-2 grid-cols-2 md:grid-cols-1 w-full">
            <div className="font-medium text-gray-800">Fundraiser Details:</div>
            <div className="border rounded-lg p-5 space-y-2">
              <div className="rounded-xl overflow-hidden h-72 aspect-[4/3] object-cover">
                <img
                  src={review?.cover_picture}
                  className="w-full h-full m-auto object-cover"
                />
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">Title: </div>
                <h3 className="font-medium text-tremor-content-strong">
                  {review?.title}
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">Category: </div>
                <h3 className="text-black">
                  {Object.keys(CategoryIndexMap).find(
                    (key) => CategoryIndexMap[key] === review?.category_id
                  )}
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">Location: </div>
                <h3 className="text-black">
                  {`${review?.province}, ${review?.country}`}
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">Link: </div>
                <a
                  target="_blank"
                  href={`http://localhost:5173/fundraiser/${review?.id}`}
                  className="flex text-black hover:underline text-sm"
                >
                  Go to Fundraiser
                  <IoOpenOutline className="ml-1 w-5 h-5" />
                </a>
              </div>
            </div>

            <div className="font-medium text-gray-800">Fundraiser Status:</div>
            <div className="border rounded-lg p-5 space-y-2">
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">
                  Accumulated amount:{" "}
                </div>
                <h3 className="text-black">
                  ₫
                  {review?.current_amount &&
                    utils.formatNumber(review?.current_amount)}
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">
                  Goal amount:{" "}
                </div>
                <h3 className="font-medium text-tremor-content-strong">
                  ₫
                  {review?.goal_amount &&
                    utils.formatNumber(review?.goal_amount)}
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">
                  No. donations:{" "}
                </div>
                <h3 className="text-black">{review?.backing_count}</h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">
                  Transfer setup:{" "}
                </div>
                {review?.card_id ? (
                  <>
                    <Badge color="green">Yes</Badge>
                  </>
                ) : (
                  <>
                    <Badge color="red">No</Badge>
                  </>
                )}
              </div>

              <div className="flex space-x-2 items-baseline">
                <div className="underline text-black text-sm">Started at: </div>
                <h3 className="text-black">
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(review?.start_date))
                  )}
                </h3>
                <div className="underline text-black text-sm">Ended at: </div>
                <h3 className="text-black">
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(review?.end_date))
                  )}
                </h3>
              </div>
            </div>
            <div className="mt-8 w-full flex space-x-1">
              <Button className="w-1/2" onClick={handlePayout}>
                Payout
              </Button>
              <Button
                className="w-1/2"
                color="red"
                onClick={() => setIsOpenReview(false)}
              >
                Dispute
              </Button>
            </div>
          </div>
        </DialogPanel>
      </Dialog>
    </>
  );
};

PendingPayoutTable.propTypes = {
  token: PropTypes.string,
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default PendingPayoutTable;
