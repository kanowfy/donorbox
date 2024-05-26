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
import { Checkbox } from "flowbite-react";
import PropTypes from "prop-types";
import { useState } from "react";
import { RiRefund2Fill } from "react-icons/ri";
import { IoIosCloseCircle } from "react-icons/io";

const PendingRefundTable = ({ data, setSuccess }) => {
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [review, setReview] = useState();
  const [isSelectedAll, setIsSelectedAll] = useState(false);
  const [selectedList, setSelectedList] = useState([]);
  return (
    <>
      <Card className="my-5 shadow-lg shadow-gray-500 space-y-2">
        <h3 className="text-tremor-content-strong dark:text-dark-tremor-content-strong font-semibold">
          List of projects eligible for refund
        </h3>

        <div className="space-y-1 pt-4">
          <Button
            color="green"
            disabled={selectedList.length == 0}
            icon={RiRefund2Fill}
            onClick={() => setSuccess(true)}
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
        <Table className="mt-5">
          <TableHead>
            <TableRow>
              <TableHeaderCell>
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
              </TableHeaderCell>
              <TableHeaderCell>ID</TableHeaderCell>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>Goal Amount</TableHeaderCell>
              <TableHeaderCell>Accumulated Amount</TableHeaderCell>
              <TableHeaderCell>End date</TableHeaderCell>
              <TableHeaderCell></TableHeaderCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.map((item) => (
              <TableRow key={item.id}>
                <TableCell>
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
                </TableCell>
                <TableCell>{item.id}</TableCell>
                <TableCell>{item.title}</TableCell>
                <TableCell>{item.goal_amount.toLocaleString()}</TableCell>
                <TableCell>
                  {item.accumulated_amount.toLocaleString()}
                </TableCell>
                <TableCell>{item.end_date}</TableCell>
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
        <DialogPanel>
          <div className="flex justify-end">
            <div
              onClick={() => setIsOpenReview(false)}
              className="hover:cursor-pointer"
            >
              <IoIosCloseCircle className="w-6 h-6" />
            </div>
          </div>
          <div className="space-y-2">
            <h3 className="text-lg font-semibold text-tremor-content-strong dark:text-dark-tremor-content-strong">
              {review?.title}
            </h3>
            <img src={review?.cover_picture} />
            <p className="mt-2 leading-6 text-tremor-default text-tremor-content dark:text-dark-tremor-content">
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Quaerat,
              nihil inventore sequi exercitationem ex impedit blanditiis
              officiis perspiciatis porro aspernatur quidem praesentium velit
              aliquid culpa dolore in nulla quod voluptatibus.
            </p>
          </div>
        </DialogPanel>
      </Dialog>
    </>
  );
};

PendingRefundTable.propTypes = {
  data: PropTypes.array,
  setSuccess: PropTypes.func,
};

export default PendingRefundTable;
