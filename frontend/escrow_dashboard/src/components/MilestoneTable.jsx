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
} from "@tremor/react";
import PropTypes from "prop-types";
import { useState } from "react";
import utils from "../utils/utils";
import { IoOpenOutline } from "react-icons/io5";
import { CategoryIndexMap } from "../constants";
import { IoIosCloseCircle } from "react-icons/io";
import escrowService from "../services/escrow";
import { Label, TextInput } from "flowbite-react";

const MilestoneTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [review, setReview] = useState();

  const handleConfirm = async () => {};

  return (
    <div>
      <Card className="my-5 shadow-lg shadow-gray-500">
        <h3 className="text-tremor-content-strong dark:text-dark-tremor-content-strong">
          List of milestones waiting to be resolved
        </h3>
        <Table className="mt-5">
          <TableHead>
            <TableRow>
              <TableHeaderCell>ID</TableHeaderCell>
              <TableHeaderCell>Project Link</TableHeaderCell>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>Fund Goal</TableHeaderCell>
              <TableHeaderCell>Total Fund</TableHeaderCell>
              <TableHeaderCell></TableHeaderCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((item) => (
              <TableRow key={item.id}>
                <TableCell>{item.id}</TableCell>
                <TableCell>
                <a
                  target="_blank"
                  href={`http://localhost:4001/fundraiser/${item?.project_id}`}
                  className="flex text-gray-700 hover:font-semibold text-sm hover:text-blue-700"
                >
                  <IoOpenOutline className="ml-1 w-5 h-5" />
                </a>
                </TableCell>
                <TableCell>{item.title}</TableCell>
                <TableCell>₫{item.fund_goal.toLocaleString()}</TableCell>
                <TableCell>₫{item.current_fund.toLocaleString()}</TableCell>
                <TableCell>
                  <Button
                    variant="secondary"
                    onClick={() => {
                      setIsOpenReview(true);
                      setReview(item);
                    }}
                  >
                    View
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
            <div className="border rounded-lg p-5 space-y-2">
              <div className="flex space-x-2 items-baseline">
                <h3 className="text-xl font-semibold text-gray-700">
                  Milestone description
                </h3>
              </div>

              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">Title: </div>
                <h3 className="text-gray-700">
                  {review?.title}
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">
                  Description:{" "}
                </div>
                <h3 className="text-gray-700">{review?.description}</h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">
                  Accumulated fund:{" "}
                </div>
                <h3 className="text-gray-700">
                  {`₫${review?.current_fund.toLocaleString()} / ₫${review?.fund_goal.toLocaleString()}`}
                </h3>
              </div>
            </div>
            <div className="border rounded-lg p-5 space-y-2">
              <div className="flex space-x-2 items-baseline">
                <h3 className="text-xl font-semibold text-gray-700">
                  Resolving milestone
                </h3>
              </div>

            <div className="flex space-x-2 items-baseline">
              <div className="font-semibold text-black text-sm">
                Bank description:{" "}
              </div>
              <h3 className="text-gray-700">{review?.bank_description}</h3>
            </div>

            </div>

            <div className="mt-8 w-full flex space-x-1">
              <Button className="w-1/2" onClick={handleConfirm}>
                Confirm resolved
              </Button>
              <Button
                className="w-1/2"
                color="gray"
                variant="secondary"
                onClick={() => setIsOpenReview(false)}
              >
                Cancel
              </Button>
            </div>
          </div>
        </DialogPanel>
      </Dialog>
    </div>
  );
};

MilestoneTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default MilestoneTable;
