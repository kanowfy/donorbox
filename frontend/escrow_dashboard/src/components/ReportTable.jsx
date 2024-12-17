import { Card } from "@tremor/react";
import PropTypes from "prop-types";
import { useState } from "react";
import utils from "../utils/utils";
import { IoOpenOutline } from "react-icons/io5";
import escrowService from "../services/escrow";
import {
  Label,
  Modal,
  Textarea,
  Timeline,
  Button,
  Badge,
} from "flowbite-react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { Table } from "flowbite-react";

const ReportStatusMap = {
  pending: {
    color: "warning",
    text: "pending",
  },
  dismissed: {
    color: "failure",
    text: "dismissed",
  },
  resolved: {
    color: "info",
    text: "resolved",
  },
};

const ReportTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const navigate = useNavigate();
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [review, setReview] = useState();

  const handleApprove = async () => {
    try {
      await escrowService.reviewReport(token, {
        report_id: Number(review?.id),
      });
      setIsSuccessful(true);
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      console.error(err);
      setIsFailed(true);
    }
  };

  const handleReject = async () => {
    try {
      await escrowService.reviewReport(token, {
        report_id: Number(review?.id),
        mark_dispute: true,
      });
      setIsSuccessful(true);
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      console.error(err);
      setIsFailed(true);
    }

  };

  return (
    <div>
      <Card className="my-5 shadow-lg shadow-gray-500">
        <h3 className="text-tremor-content-strong dark:text-dark-tremor-content-strong">
          List of project reports
        </h3>
        <Table className="mt-5" striped hoverable>
          <Table.Head className="w-fit">
            <Table.HeadCell>ID</Table.HeadCell>
            <Table.HeadCell>Fundraiser Link</Table.HeadCell>
            <Table.HeadCell>Reporter Name</Table.HeadCell>
            <Table.HeadCell>Reason</Table.HeadCell>
            <Table.HeadCell>Submission Date</Table.HeadCell>
            <Table.HeadCell>Status</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {data?.map((item) => (
              <Table.Row key={item.id}>
                <Table.Cell>{item.id}</Table.Cell>
                <Table.Cell>
                  <a
                    target="_blank"
                    href={`http://localhost:4001/fundraiser/${item?.project_id}`}
                    className="flex text-gray-700 hover:font-semibold text-sm hover:text-blue-700"
                  >
                    <IoOpenOutline className="ml-1 w-5 h-5" />
                  </a>
                </Table.Cell>
                <Table.Cell>{item.full_name}</Table.Cell>
                <Table.Cell>{item.reason}</Table.Cell>
                <Table.Cell>
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(item.created_at))
                  )}
                </Table.Cell>
                <Table.Cell>
                  <Badge color={ReportStatusMap[item.status].color} className="w-fit">
                    {ReportStatusMap[item.status].text}
                  </Badge>
                </Table.Cell>
                <Table.Cell>
                  <button
                    className="font-semibold hover:underline text-cyan-600"
                    onClick={() => {
                      setIsOpenReview(true);
                      setReview(item);
                    }}
                  >
                    Review
                  </button>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Card>
      <Modal
        show={isOpenReview}
        onClose={() => setIsOpenReview(false)}
        size="2xl"
      >
        <Modal.Header>Review Project Report</Modal.Header>
        <Modal.Body>
          <div className="space-y-2 grid-cols-2 md:grid-cols-1 w-full">
            <div className="border rounded-lg p-5 space-y-2">
              <div className="flex space-x-2 items-baseline">
                <h3 className="text-xl font-semibold text-blue-700 underline">
                  Report description
                </h3>
              </div>

              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">Reason: </div>
                <h3 className="text-red-700">{review?.reason}</h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">
                  Reporter relation with fundraiser:{" "}
                </div>
                <h3 className="text-blue-700">
                  {review?.relation ? review.relation : "None"}
                </h3>
              </div>
              <div>
                <div className="font-semibold text-black text-sm">
                  Details:{" "}
                  <span className="text-gray-700 text-base font-normal">{review?.details}</span>
                </div>
              </div>
            </div>
            <div className="border rounded-lg p-5 space-y-2">
              <div className="flex space-x-2 items-baseline">
                <h3 className="text-xl font-semibold text-blue-700 underline">
                  Reporter Contacts
                </h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">
                  Full Name:{" "}
                </div>
                <h3 className="text-gray-700">{review?.full_name}</h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">Email: </div>
                <h3 className="text-gray-700">{review?.email}</h3>
              </div>
              <div className="flex space-x-2 items-baseline">
                <div className="font-semibold text-black text-sm">
                  Phone Number:{" "}
                </div>
                <h3 className="text-gray-700">{review?.phone_number}</h3>
              </div>
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="w-full flex justify-center">
            <div className="w-1/2 flex space-x-1">
              <Button
                className="w-1/2 font-semibold"
                onClick={handleApprove}
                color="blue"
                disabled={review?.status !== "pending"}
              >
                Dismiss
              </Button>
              <Button
                className="w-1/2 font-semibold"
                color="failure"
                onClick={handleReject}
                disabled={review?.status !== "pending"}
              >
                Mark Disputed
              </Button>
            </div>
          </div>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

ReportTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default ReportTable;
