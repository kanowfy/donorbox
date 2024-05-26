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

const PendingPayoutTable = ({ data, setSuccess }) => {
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [review, setReview] = useState();
  return (
    <>
      <Card className="my-5 shadow-lg shadow-gray-500">
        <h3 className="text-tremor-content-strong dark:text-dark-tremor-content-strong font-semibold">
          List of projects eligible for payout
        </h3>
        <Table className="mt-5">
          <TableHead>
            <TableRow>
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
        <DialogPanel>
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
            <div className="mt-8 w-full flex space-x-1">
              <Button
                className="w-1/2"
                onClick={() => {
                  setIsOpenReview(false);
                  setSuccess(true);
                  setTimeout(() => {
                    setSuccess(false);
                  }, 1500);
                }}
              >
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
  data: PropTypes.array,
  setSuccess: PropTypes.func,
};

export default PendingPayoutTable;
