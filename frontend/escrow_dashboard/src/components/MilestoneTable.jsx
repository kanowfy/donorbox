import PropTypes from "prop-types";
import { useState, useEffect } from "react";
import { IoOpenOutline } from "react-icons/io5";
import escrowService from "../services/escrow";
import { Button, FileInput, Modal, Table } from "flowbite-react";
import Cleave from "cleave.js/react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { CLIENT_URL } from "../constants";
import utils from "../utils/utils";

const MilestoneTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const navigate = useNavigate();
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [isOpenConfirm, setIsOpenConfirm] = useState(false);
  const [image, setImage] = useState();
  const [review, setReview] = useState();
  const [isLoading, setIsLoading] = useState(false);
  const [description, setDescription] = useState();

  /*
  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm();
   */

  const onSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);
    try {
      const payload = {
        amount: Number(review?.milestone.current_fund),
        description: description,
      };

      if (image) {
        payload.image = await utils.uploadImage(image);
      }

      await escrowService.resolveMilestone(token, review?.milestone.id, payload);
      setIsLoading(false);
      setIsSuccessful(true);
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      setIsFailed(true);
      console.error(err);
    }
  };

  const onSelectImage = (e) => {
    if (!e.target.files || e.target.files.length == 0) {
      setImage(undefined);
      return;
    }

    setImage(e.target.files);
  };

  return (
    <div>
        <Table className="mt-5" hoverable striped>
          <Table.Head>
              <Table.HeadCell>ID</Table.HeadCell>
              <Table.HeadCell>Project Link</Table.HeadCell>
              <Table.HeadCell>Title</Table.HeadCell>
              <Table.HeadCell>Fund Goal</Table.HeadCell>
              <Table.HeadCell>Total Fund</Table.HeadCell>
              <Table.HeadCell></Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {data?.map((item) => (
              <Table.Row key={item.milestone.id}>
                <Table.Cell>{item.milestone.id}</Table.Cell>
                <Table.Cell>
                  <a
                    target="_blank"
                    href={`${CLIENT_URL}/fundraiser/${item?.milestone.project_id}`}
                    className="flex text-gray-700 hover:font-semibold text-sm hover:text-blue-700"
                  >
                    <IoOpenOutline className="ml-1 w-5 h-5" />
                  </a>
                </Table.Cell>
                <Table.Cell>{item.milestone.title}</Table.Cell>
                <Table.Cell>
                  ₫{item.milestone.fund_goal.toLocaleString()}
                </Table.Cell>
                <Table.Cell>
                  ₫{item.milestone.current_fund.toLocaleString()}
                </Table.Cell>
                <Table.Cell>
                  <button
                  className="font-semibold text-cyan-600 hover:underline"
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
      <Modal
        show={isOpenReview}
        onClose={() => setIsOpenReview(false)}
        size={`${isOpenConfirm ? "5xl" : "xl"}`}
      >
        <Modal.Body>
          <div
            className={`grid ${
              isOpenConfirm ? "grid-cols-2" : "grid-cols-1"
            } gap-2`}
          >
              <form
                id="form1"
                className={`space-y-4 ${isOpenConfirm ? "" : "hidden"} border rounded-lg p-4`}
                onSubmit={(e) => onSubmit(e)}
              >
                <div>
                  <label className="block mb-1 font-medium text-gray-900">
                    Receipt photo <span className="text-red-600">*</span>
                  </label>
                  <div className="flex flex-col items-start space-y-4">
                    <FileInput
                      accept="image/png, image/jpeg"
                      color="gray"
                      sizing="lg"
                      onChange={onSelectImage}
                      multiple
                    />
                  </div>
                </div>
                <div>
                  <label className="block mb-1 font-medium text-gray-900">
                    Note
                  </label>
                  <textarea
                    onChange={e => setDescription(e.target.value)}
                    value={description}
                    rows={3}
                    placeholder=""
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                  />
                </div>
              </form>
            <div className="space-y-2 grid-cols-2 md:grid-cols-1 w-full">
              <div className="border rounded-lg p-5 space-y-2">
                <div className="flex space-x-2 items-baseline">
                  <h3 className="text-xl font-semibold text-blue-700 underline">
                    Milestone description
                  </h3>
                </div>

                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Title:{" "}
                  </div>
                  <h3 className="text-gray-700">{review?.milestone.title}</h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Description:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {review?.milestone.description}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Accumulated fund:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {`₫${review?.milestone.current_fund.toLocaleString()} / ₫${review?.milestone.fund_goal.toLocaleString()}`}
                  </h3>
                </div>
              </div>
              <div className="border rounded-lg p-5 space-y-2">
                <div className="flex space-x-2 items-baseline">
                  <h3 className="text-xl font-semibold text-blue-700 underline">
                    Beneficiary Information
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">Name: </div>
                  <h3 className="text-gray-700">{review?.receiver_name}</h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Number:{" "}
                  </div>
                  <h3 className="text-gray-700">{review?.receiver_number}</h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Address:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {`${review?.address}, ${review?.district}, ${review?.city}, ${review?.country}`}
                  </h3>
                </div>
              </div>
              <div className="border rounded-lg p-5 space-y-2">
                <div className="flex space-x-2 items-baseline">
                  <h3 className="text-xl font-semibold text-blue-700 underline">
                    Resolving milestone
                  </h3>
                </div>

                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Bank description:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {review?.milestone.bank_description}
                  </h3>
                </div>
              </div>
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="w-full flex space-x-2">
            {!isOpenConfirm && (
              <Button
                className="w-1/2"
                onClick={() => setIsOpenConfirm(true)}
                color="blue"
                type="button"
              >
                Confirm completion
              </Button>
            )}
            {isOpenConfirm && (
              <Button
                className="w-1/2"
                type="submit"
                form="form1"
                isProcessing={isLoading}
                color="blue"
              >
                Submit Completion
              </Button>

            )}            
            <Button
              className="w-1/2"
              color="gray"
              variant="secondary"
              onClick={() => {
                setIsOpenReview(false);
                setIsOpenConfirm(false);
                setImage(undefined);
              }}
            >
              Cancel
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

MilestoneTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default MilestoneTable;
