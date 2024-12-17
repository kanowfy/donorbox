import PropTypes from "prop-types";
import { useState, useEffect } from "react";
import { IoOpenOutline } from "react-icons/io5";
import utils from "../utils/utils";
import escrowService from "../services/escrow";
import { Button, Modal, Table, Label, Textarea } from "flowbite-react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

const MilestoneProofTable = ({ token, data, setIsSuccessful, setIsFailed }) => {
  const navigate = useNavigate();
  const [isOpenReview, setIsOpenReview] = useState(false);
  const [isOpenReject, setIsOpenReject] = useState(false);
  const [review, setReview] = useState();
  const [isLoading, setIsLoading] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    reset
  } = useForm();

  const handleApprove = async () => {
    setIsLoading(true);
    try {
      await escrowService.reviewProof(token, {
        proof_id: Number(getPendingProof(review?.milestone.spending_proofs)["id"]),
        approved: true,
      });
      setIsSuccessful(true);
      setIsLoading(false);
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      setIsLoading(false);
      console.error(err);
      setIsFailed(true);
    }
  };

  const handleReject = async (data) => {
    try {
      await escrowService.reviewProof(token, {
        proof_id: Number(getPendingProof(review?.milestone.spending_proofs)["id"]),
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

  const getPendingProof = (proofs) => {
    return proofs?.filter(p => p.status === "pending")[0];
  }

  return (
    <div>
      <Table className="mt-5">
        <Table.Head striped hoverable>
          <Table.HeadCell>ID</Table.HeadCell>
          <Table.HeadCell>Project Link</Table.HeadCell>
          <Table.HeadCell>Title</Table.HeadCell>
          <Table.HeadCell>Fund Released On</Table.HeadCell>
          <Table.HeadCell>Proof Submitted On</Table.HeadCell>
          <Table.HeadCell></Table.HeadCell>
        </Table.Head>
        <Table.Body className="divide-y">
          {data?.map((item) => (
            <Table.Row key={item.milestone.id}>
              <Table.Cell>{item.milestone.id}</Table.Cell>
              <Table.Cell>
                <a
                  target="_blank"
                  href={`http://localhost:4001/fundraiser/${item?.milestone.project_id}`}
                  className="flex text-gray-700 hover:font-semibold text-sm hover:text-blue-700"
                >
                  <IoOpenOutline className="ml-1 w-5 h-5" />
                </a>
              </Table.Cell>
              <Table.Cell>{item.milestone.title}</Table.Cell>
              <Table.Cell>
                {utils.formatDate(
                  new Date(
                    utils.parseDateFromRFC3339(
                      item?.milestone.milestone_completion.created_at
                    )
                  )
                )}
              </Table.Cell>
              <Table.Cell>
                {utils.formatDate(
                  new Date(
                    utils.parseDateFromRFC3339(
                      getPendingProof(item?.milestone.spending_proofs)["created_at"]
                    )
                  )
                )}
              </Table.Cell>
              <Table.Cell>
                <button
                  className="font-semibold text-cyan-600 hover:underline"
                  onClick={() => {
                    setIsOpenReview(true);
                    setReview(item);
                  }}
                >
                  Review Proof
                </button>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
      { review && (<Modal
        show={isOpenReview}
        onClose={() => setIsOpenReview(false)}
        size="3xl"
      >
        <Modal.Header>Review Spending Proof</Modal.Header>
        <Modal.Body>
          <div>
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
              </div>
              <div className="border rounded-lg p-5 space-y-2">
                <div className="flex space-x-2 items-baseline">
                  <h3 className="text-xl font-semibold text-blue-700 underline">
                    Spending Proof
                  </h3>
                </div>
                <div className="flex justify-center">
                  <div className="grid grid-cols-5 gap-12 px-5">
                    <div className="col-span-2">
                      <div className="font-semibold text-black text-sm">
                        Transfer Receipt:{" "}
                      </div>
                      <div className="rounded-xl overflow-hidden h-96 aspect-[3/5] object-cover">
                        <img
                          src={getPendingProof(review?.milestone.spending_proofs)["transfer_image"]}
                          className="w-full h-full m-auto object-cover"
                        />
                      </div>
                    </div>
                    <div className="col-span-3 h-full">
                      <div className="font-semibold text-black text-sm">
                        Media Proof:{" "}
                      </div>
                      <div className="flex justify-center flex-col h-full">
                        <div className="rounded-xl overflow-hidden h-52 aspect-[16/9] object-cover">
                          <img
                          src={getPendingProof(review?.milestone.spending_proofs)["proof_media"]}
                            className="w-full h-full m-auto object-cover"
                          />
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Description:{" "}
                  </div>
                  <h3 className="text-gray-700">
                    {getPendingProof(review?.milestone.spending_proofs)["description"]}
                  </h3>
                </div>
                <div className="flex space-x-2 items-baseline">
                  <div className="font-semibold text-black text-sm">
                    Submitted on:{" "}
                  </div>
                  <h3 className="text-gray-800">
                    {utils.formatDate(
                      new Date(
                        utils.parseDateFromRFC3339(getPendingProof(review?.milestone.spending_proofs)["created_at"])
                      )
                    )}
                  </h3>
                </div>
              </div>
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="w-full flex space-x-2">
            <Button
              onClick={handleApprove}
              className="w-1/2 font-bold"
              isProcessing={isLoading}
              color="blue"
            >
              Confirm
            </Button>
            <Button
              className="w-1/2 font-bold"
              color="failure"
              onClick={() => setIsOpenReject(true)}
            >
              Refute
            </Button>
          </div>
        </Modal.Footer>
      </Modal>)}
      <Modal
        show={isOpenReject}
        onClose={() => {
          setIsOpenReject(false);
          reset();
        }}
        size="md"
        popup
      >
        <Modal.Header />
        <Modal.Body>
          <form
            className="flex max-w-sm flex-col space-y-2"
            onSubmit={handleSubmit(handleReject)}
          >
            <div className="mbblock">
              <Label
                className="font-semibold"
                htmlFor="reason"
                value="Provide reason for refuting the milestone spending proof:"
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
            <Button type="submit" color="info">
              Submit
            </Button>
          </form>
        </Modal.Body>
      </Modal>
    </div>
  );
};

MilestoneProofTable.propTypes = {
  data: PropTypes.array,
  setIsSuccessful: PropTypes.func,
  setIsFailed: PropTypes.func,
};

export default MilestoneProofTable;
