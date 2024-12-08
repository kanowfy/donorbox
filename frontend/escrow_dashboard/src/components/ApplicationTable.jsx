import {
  Table,
  TableHead,
  TableHeaderCell,
  TableBody,
  TableRow,
  TableCell,
  Card,
  Button,
} from "@tremor/react";
import PropTypes from "prop-types";
import { useState } from "react";
import utils from "../utils/utils";
import { IoOpenOutline } from "react-icons/io5";
import { CategoryIndexMap } from "../constants";
import escrowService from "../services/escrow";
import {
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Textarea,
  Timeline,
} from "flowbite-react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

const ApplicationTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const navigate = useNavigate();
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [isOpenReject, setIsOpenReject] = useState(false);
  const [review, setReview] = useState();
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm();

  const handleApprove = async () => {
    try {
      await escrowService.reviewProject(token, {
        project_id: Number(review?.project.id),
        approved: true,
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

  const handleReject = async (data) => {
    try {
      await escrowService.reviewProject(token, {
        project_id: Number(review?.project.id),
        reject_reason: data.reason,
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
          List of projects waiting for approval
        </h3>
        <Table className="mt-5">
          <TableHead>
            <TableRow>
              <TableHeaderCell>ID</TableHeaderCell>
              <TableHeaderCell>Title</TableHeaderCell>
              <TableHeaderCell>Fund Goal</TableHeaderCell>
              <TableHeaderCell>End date</TableHeaderCell>
              <TableHeaderCell></TableHeaderCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data?.map((item) => (
              <TableRow key={item.project.id}>
                <TableCell>{item.project.id}</TableCell>
                <TableCell>{item.project.title}</TableCell>
                <TableCell>
                  ₫{item.project.fund_goal.toLocaleString()}
                </TableCell>
                <TableCell>
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(item.project.end_date))
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
      <Modal
        show={isOpenReview}
        onClose={() => setIsOpenReview(false)}
        size="7xl"
      >
        <Modal.Header>Review Project Application</Modal.Header>
        <Modal.Body>
          <div className="grid grid-cols-3 space-x-4">
            <div className="space-y-2 grid-cols-2 md:grid-cols-1 w-full">
              <div className="border rounded-lg p-5 space-y-2 overflow-hidden">
                <div className="font-bold text-lg text-gray-700 underline">Basic Information</div>
                <div className="rounded-xl overflow-hidden max-h-fit aspect-[5/3] object-cover mx-auto">
                  <img
                    src={review?.project.cover_picture}
                    className="w-full h-full m-auto object-cover"
                  />
                </div>
                <div className="flex space-x-2 items-baseline">
                  {/*<div className="font-semibold text-black text-sm">Title: </div>*/}
                  <h3 className="text-xl font-semibold text-gray-700">
                    {review?.project.title}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Category:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {Object.keys(CategoryIndexMap).find(
                      (key) =>
                        CategoryIndexMap[key] === review?.project.category_id
                    )}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Started at:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {utils.formatDate(
                      new Date(
                        utils.parseDateFromRFC3339(review?.project.created_at)
                      )
                    )}
                  </h3>
                  <div className="font-semibold text-black text-sm">
                    Ends at:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {utils.formatDate(
                      new Date(
                        utils.parseDateFromRFC3339(review?.project.end_date)
                      )
                    )}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">Link: </div>
                  <a
                    target="_blank"
                    href={`http://localhost:4001/fundraiser/${review?.project.id}`}
                    className="flex text-gray-700 hover:font-semibold text-sm hover:text-blue-700 hover:underline"
                  >
                    Go to Fundraiser
                    <IoOpenOutline className="ml-1 w-5 h-5" />
                  </a>
                </div>
              </div>
            </div>
            <div className="border rounded-lg p-5 space-y-2">
              <div className="font-bold text-lg text-gray-700 underline">Milestones</div>
              <Timeline>
                {review?.milestones.map((m) => (
                  <Timeline.Item>
                    <Timeline.Point />
                    <Timeline.Content>
                      <Timeline.Title>{m.title}</Timeline.Title>
                      <Timeline.Content>
                        <div>
                        <div className="text-blue-600 font-semibold">Fund goal:</div>
                        <div>{`₫${m.fund_goal.toLocaleString()}`}</div>
                        </div>
                        <div>
                        <div className="text-blue-600 font-semibold">Bank description:</div>
                        <div>{m.bank_description}</div>
                        </div>
                      </Timeline.Content>
                    </Timeline.Content>
                  </Timeline.Item>
                ))}
              </Timeline>
            </div>
            <div className="border rounded-lg p-5 space-y-1">
              <div className="font-bold text-lg text-gray-700 underline">Beneficiary Information</div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-blue-600 text-sm">
                    Name:{" "}
                  </div>
                  <h3>
                    {review?.project.receiver_name}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-blue-600 text-sm">
                    Phone number:{" "}
                  </div>
                  <h3>
                    {review?.project.receiver_number}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-blue-600 text-sm">
                    Address:{" "}
                  </div>
                  <h3>
                    {`${review?.project.address}, ${review?.project.district}, ${review?.project.city}, ${review?.project.country}`}
                  </h3>
                </div>
              </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="w-full flex justify-center">
          <div className="w-1/2 flex space-x-1">
            <Button className="w-1/2" onClick={handleApprove} size="lg">
              Approve
            </Button>
            <Button
              className="w-1/2"
              color="red"
              size="lg"
              onClick={() => setIsOpenReject(true)}
            >
              Reject
            </Button>
          </div>

          </div>
        </Modal.Footer>
      </Modal>
      <Modal
        show={isOpenReject}
        onClose={() => {
          setIsOpenReject(false);
          reset();
        }}
        size="md"
        popup
      >
        <ModalHeader />
        <ModalBody>
          <form
            className="flex max-w-sm flex-col space-y-2"
            onSubmit={handleSubmit(handleReject)}
          >
            <div className="mbblock">
              <Label
                htmlFor="reason"
                value="Provide reason why this fundraiser is ineligible:"
              />
            </div>
            <Textarea
              {...register("reason", {
                required: "Please provide a reason for rejection",
              })}
              id="reason"
              type="text"
              placeholder=""
              rows={4}
            />
            {errors.reason?.type === "required" && (
              <p className="text-red-600 text-sm">{errors.reason.message}</p>
            )}
            <Button type="submit" color="teal">
              Submit
            </Button>
          </form>
        </ModalBody>
      </Modal>
    </div>
  );
};

ApplicationTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default ApplicationTable;
