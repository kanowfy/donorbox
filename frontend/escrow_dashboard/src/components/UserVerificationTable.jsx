import {
  Card,
} from "@tremor/react";
import PropTypes from "prop-types";
import { useState } from "react";
import {
  Label,
  Modal,
  ModalBody,
  ModalHeader,
  Textarea,
  Button,
  Dropdown,
  Table
} from "flowbite-react";
import { useForm } from "react-hook-form";
import utils from "../utils/utils";
import escrowService from "../services/escrow";
import { IoOpenOutline } from "react-icons/io5";
import { IoIosMore } from "react-icons/io";
import { useNavigate } from "react-router-dom";

const UserVerificationTable = ({
  token,
  data,
  setIsSuccessful,
  setIsFailed,
}) => {
  const navigate = useNavigate(0);
  const [isOpenReject, setIsOpenReject] = useState(false);
  const [reviewID, setReviewID] = useState();
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm();

  const handleApprove = async (reviewID) => {
    try {
      await escrowService.reviewVerification(token, {
        user_id: Number(reviewID),
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
      await escrowService.reviewVerification(token, {
        user_id: Number(reviewID),
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
          List of pending verification documents
        </h3>
        <Table className="mt-5">
          <Table.Head>
              <Table.HeadCell>User ID</Table.HeadCell>
              <Table.HeadCell>Email</Table.HeadCell>
              <Table.HeadCell>Name</Table.HeadCell>
              <Table.HeadCell>Verification Document</Table.HeadCell>
              <Table.HeadCell>Account Created</Table.HeadCell>
              <Table.HeadCell>Action</Table.HeadCell>
          </Table.Head>
          <Table.Body>
            {data?.map((item) => (
              <Table.Row key={item.id}>
                <Table.Cell className="text-gray-600">{item.id}</Table.Cell>
                <Table.Cell className="text-gray-600">{item.email}</Table.Cell>
                <Table.Cell className="text-gray-600">{`${item.first_name} ${item?.last_name}`}</Table.Cell>
                <Table.Cell>
                  <a
                    target="_blank"
                    href={item.document_url}
                    className="text-blue-700 hover:underline"
                  >
                    <div className="flex">
                      <div>View Document</div>
                      <IoOpenOutline className="w-5 h-5 ml-1" />
                    </div>
                  </a>
                </Table.Cell>
                <Table.Cell className="text-gray-600">
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(item.created_at))
                  )}
                </Table.Cell>
                <Table.Cell>
                  <Dropdown
                    label=""
                    dismissOnClick={false}
                    renderTrigger={() => (
                      <span className="cursor-pointer">
                        <IoIosMore className="w-7 h-7" />
                      </span>
                    )}
                  >
                    <Dropdown.Item>
                      <span
                        className="text-green-600 font-semibold"
                        onClick={() => handleApprove(item.id)}
                      >
                        APPROVE
                      </span>
                    </Dropdown.Item>
                    <Dropdown.Item>
                      <span
                        className="text-red-500 font-semibold"
                        onClick={() => {
                          setIsOpenReject(true);
                          setReviewID(item.id);
                        }}
                      >
                        REJECT
                      </span>
                    </Dropdown.Item>
                  </Dropdown>
                </Table.Cell>
              </Table.Row>
            ))}
          </Table.Body>
        </Table>
      </Card>
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
            <div className="block">
              <Label
                htmlFor="reason"
                value="Provide reason why this document is invalid:"
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

UserVerificationTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default UserVerificationTable;
