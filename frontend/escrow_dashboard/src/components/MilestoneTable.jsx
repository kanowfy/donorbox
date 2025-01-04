import PropTypes from "prop-types";
import { useState, useEffect } from "react";
import { IoOpenOutline } from "react-icons/io5";
import escrowService from "../services/escrow";
import uploadService from "../services/upload";
import { Button, FileInput, Modal, Table } from "flowbite-react";
import Cleave from "cleave.js/react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { CLIENT_URL } from "../constants";

const MilestoneTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const navigate = useNavigate();
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [isOpenConfirm, setIsOpenConfirm] = useState(false);
  const [image, setImage] = useState();
  const [preview, setPreview] = useState();
  const [review, setReview] = useState();
  const [isLoading, setIsLoading] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
  } = useForm();

  useEffect(() => {
    if (!image) {
      setPreview(undefined);
      return;
    }

    const objectUrl = URL.createObjectURL(image);
    setPreview(objectUrl);

    return () => URL.revokeObjectURL(image);
  }, [image]);

  const onSubmit = async (data) => {
    console.log(data, "id", review?.milestone.id);
    setIsLoading(true);
    try {
      const payload = {
        amount: Number(data.amount),
        description: data.description,
      };

      if (image) {
        payload.image = await uploadImage(image);
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

    setImage(e.target.files[0]);
  };

  const uploadImage = async (image) => {
    if (!image) {
      throw new Error("Missing image");
    }

    const formData = new FormData();
    formData.append("file", image);

    const response = await uploadService.uploadImage(formData);
    return response.url;
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
                onSubmit={handleSubmit(onSubmit)}
              >
                <div className="flex items-baseline space-x-1 mt-5">
                  <label className="block mb-2 font-medium text-gray-900">
                    Amount transferred <span className="text-red-600">*</span>
                  </label>
                  <div className="flex items-baseline bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 px-3 py-1">
                    <span className="block mb-1 font-medium">₫</span>
                    <Cleave
                      options={{
                        numeral: true,
                        numericOnly: true,
                        numeralThousandsGroupStyle: "thousand",
                        numeralPositiveOnly: true,
                      }}
                      {...register(`amount`, {
                        required: "Transferred amount is required",
                        min: 50000,
                        max: 100000000,
                      })}
                      onChange={(e) =>
                        setValue(`amount`, e.target.value.replace(/,/g, ""))
                      }
                      className="border-0 focus:ring-0 focus:border-0 bg-gray-50 autofill:bg-gray-50"
                      placeholder=""
                    />
                  </div>
                </div>
                  {errors.amount?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errors.amount?.message}
                    </p>
                  )}
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
                    />
                    {image && (
                      <div className="rounded-xl overflow-hidden h-40 aspect-[3/5] object-cover">
                        <img
                          src={preview}
                          className="w-full h-full m-auto object-cover"
                        />
                      </div>
                    )}
                  </div>
                </div>
                <div>
                  <label className="block mb-1 font-medium text-gray-900">
                    Note
                  </label>
                  <textarea
                    {...register(`description`, {})}
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
            {!isOpenConfirm ? (
              <Button
                className="w-1/2"
                onClick={() => setIsOpenConfirm(true)}
                color="blue"
              >
                Confirm completion
              </Button>
            ) : (
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
