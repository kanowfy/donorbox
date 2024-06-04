import { Button, Dialog, DialogPanel, Badge } from "@tremor/react";
import { Table } from "flowbite-react";
import { Checkbox } from "flowbite-react";
import PropTypes from "prop-types";
import { useState } from "react";
import { RiRefund2Fill } from "react-icons/ri";
import { IoIosCloseCircle } from "react-icons/io";
import fundService from "../../services/fund";
import { CategoryIndexMap } from "../../constants";
import { IoOpenOutline } from "react-icons/io5";
import utils from "../../utils/utils";
import { useNavigate } from "react-router-dom";

const PendingRefundTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [review, setReview] = useState();
  const [isSelectedAll, setIsSelectedAll] = useState(false);
  const [selectedList, setSelectedList] = useState([]);
  const navigate = useNavigate();

  const handleRefund = async () => {
    try {
      for (const id of selectedList) {
        await fundService.refund(token, id);
      }

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
      <div className="mb-5 space-y-2">
        <h3 className="font-semibold text-lg flex justify-center">
          Fundraisers eligible for refund
        </h3>

        <div className="flex items-baseline space-x-1">
          <Button
            color="green"
            disabled={selectedList.length == 0}
            icon={RiRefund2Fill}
            onClick={handleRefund}
          >
            Refund
          </Button>
          <div className="h-1">
            {selectedList.length > 0 && (
              <div className="text-tremor-content text-sm">
                {selectedList.length} row(s) selected.
              </div>
            )}
          </div>
        </div>
      </div>
      <div className="border rounded-lg shadow-lg">
        <Table hoverable>
          <Table.Head>
            <Table.HeadCell>
              <Checkbox
                color="blue"
                checked={isSelectedAll}
                onChange={() => {
                  if (isSelectedAll) {
                    setSelectedList([]);
                  } else {
                    setSelectedList(data.map((i) => i.id));
                  }
                  setIsSelectedAll(!isSelectedAll);
                }}
              />
            </Table.HeadCell>
            <Table.HeadCell>Title</Table.HeadCell>
            <Table.HeadCell>Goal Amount</Table.HeadCell>
            <Table.HeadCell>Donated Amount</Table.HeadCell>
            <Table.HeadCell># Donations</Table.HeadCell>
            <Table.HeadCell>End date</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {data?.map((item) => (
              <Table.Row key={item.id}>
                <Table.Cell>
                  <Checkbox
                    color="blue"
                    checked={selectedList.includes(item.id)}
                    onChange={() => {
                      if (selectedList.includes(item.id)) {
                        setSelectedList(() =>
                          selectedList.filter((id) => id != item.id)
                        );
                      } else {
                        setSelectedList(() => [...selectedList, item.id]);
                      }
                    }}
                  />
                </Table.Cell>
                <Table.Cell className="text-gray-900 font-medium">
                  {item.title}
                </Table.Cell>
                <Table.Cell className="text-black">
                  ₫{utils.formatNumber(item.goal_amount)}
                </Table.Cell>
                <Table.Cell className="text-black">
                  ₫{utils.formatNumber(item.current_amount)}
                </Table.Cell>
                <Table.Cell>
                  <Badge>{item.backing_count}</Badge>
                </Table.Cell>
                <Table.Cell>
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(item.end_date))
                  )}
                </Table.Cell>
                <Table.Cell>
                  <button
                    className="font-medium text-blue-600 hover:underline"
                    onClick={() => {
                      setIsOpenReview(true);
                      setReview(item);
                    }}
                  >
                    View
                  </button>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </div>
      <Dialog
        open={isOpenReview}
        onClose={(val) => setIsOpenReview(val)}
        static={true}
      >
        <DialogPanel className="w-full">
          <div
            className="flex justify-end hover:cursor-pointer"
            onClick={() => setIsOpenReview(false)}
          >
            <IoIosCloseCircle className="w-7 h-7" />
          </div>
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
          </div>
        </DialogPanel>
      </Dialog>
    </>
  );
};

PendingRefundTable.propTypes = {
  token: PropTypes.string,
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default PendingRefundTable;
