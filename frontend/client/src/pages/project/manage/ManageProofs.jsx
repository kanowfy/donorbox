import {
  Button,
  FileInput,
  Modal,
  Textarea,
  Select,
  Card,
  Label,
  Tooltip,
  Badge,
} from "flowbite-react";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import uploadService from "../../../services/upload";
import projectService from "../../../services/project";
import { useNavigate, useOutletContext } from "react-router-dom";
import { useAuthContext } from "../../../context/AuthContext";
import utils from "../../../utils/utils";

const ProofStatusMap = {
  pending: {
    tooltip: "Your proof of expenditure is being reviewed",
    color: "warning",
  },
  approved: {
    tooltip: "Your proof of expenditure is has been approved",
    color: "success",
  },
  rejected: {
    tooltip: "Your proof of expenditure is has been rejected",
    color: "failure",
  },
};

const ManageProofs = () => {
  const { token, user } = useAuthContext();
  const { project, milestones } = useOutletContext();
  const navigate = useNavigate();
  const [receipt, setReceipt] = useState();
  const [media, setMedia] = useState();
  const [preview, setPreview] = useState();
  const [releasedMilestones, setReleasedMilestones] = useState();
  const [proofedMilestones, setProofedMilestones] = useState();
  const [writeOpen, setWriteOpen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  useEffect(() => {
    setReleasedMilestones(
      milestones?.filter((m) => m.status === "fund_released")
    );

    const ms = milestones
      ?.filter((m) => m.spending_proofs && m.spending_proofs.length > 0)
      .flatMap((m) =>
        m.spending_proofs.map((p) => ({
          ...p,
          milestone_title: m.title,
        }))
      )
      .sort(
        (p1, p2) =>
          new Date(utils.parseDateFromRFC3339(p2.created_at)) -
          new Date(utils.parseDateFromRFC3339(p1.created_at))
      );
    console.log("flat map", ms);
    setProofedMilestones(ms);
  }, [milestones]);

  useEffect(() => {
    if (!receipt) {
      setPreview(undefined);
      return;
    }

    const objectUrl = URL.createObjectURL(receipt);
    setPreview(objectUrl);

    return () => URL.revokeObjectURL(receipt);
  }, [receipt]);

  const onSelectReceipt = (e) => {
    if (!e.target.files || e.target.files.length == 0) {
      setReceipt(undefined);
      return;
    }
    console.log(e.target.files);
    setReceipt(e.target.files[0]);
  };

  const onSelectMedia = (e) => {
    if (!e.target.files || e.target.files.length == 0) {
      setMedia(undefined);
      return;
    }
    console.log(e.target.files);
    setMedia(e.target.files[0]);
  };

  const onSubmit = async (data) => {
    setIsLoading(true);
    try {
      const payload = {
        milestone_id: Number(data.milestone_id),
        description: data.description,
      };

      payload.receipt = await uploadFile(receipt);
      payload.media = await uploadFile(media);

      await projectService.createProof(token, payload);
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

  const uploadFile = async (file) => {
    if (!file) {
      throw new Error("Missing file");
    }

    const formData = new FormData();
    formData.append("file", file);

    const response = await uploadService.uploadImage(formData);
    return response.url;
  };

  const validMilestoneToSubmit = (m) => {
    if (m.spending_proofs == undefined) {
      return true;
    }

    if (m.spending_proofs?.filter((s) => s.status == "pending").length == 0) {
      return true;
    }

    return false;
  };

  const openSubmission = (ms) => {
    if (!ms || ms?.length == 0) {
      return false;
    }

    let count = 0;
    for (const m of ms) {
      if (m.spending_proofs?.filter((s) => s.status == "pending" || s.status == "approved").length > 0) {
        count++;
      }
    }

    return count != ms.length;
  };

  return (
    <div>
      <div className="text-3xl font-semibold">Proofs</div>
      <div className="text-gray-600 text-sm mt-1">
        Upload proof of spending for milestones with released funds
      </div>

      {openSubmission(releasedMilestones) && (
        <div
          className={`border rounded-lg hover:cursor-text p-4 text-sm my-5 text-gray-600 hover:text-gray-800 hover:bg-gray-100 hover:border-gray-800 ${
            writeOpen ? "hidden" : ""
          }`}
          onClick={() => setWriteOpen(true)}
        >
          Upload Proof
        </div>
      )}
      <form
        className={`my-5 space-y-3 ${writeOpen ? "" : "hidden"}`}
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="max-w-sm">
          <div className="mb-2 block">
            <Label
              htmlFor="milestones"
              value="Select milestone"
              className="font-bold"
            />
          </div>
          <Select
            {...register("milestone_id", {
              required: "Milestone ID is required",
            })}
            id="milestoneOption"
            required
          >
            {releasedMilestones?.map((m) => {
              if (validMilestoneToSubmit(m)) {
                return (
                  <option key={m.id} value={m.id}>
                    {m.title}
                  </option>
                );
              }
            })}
          </Select>
        </div>

        <div className="flex gap-4">
          <div>
            <Label className="block text-sm font-bold mb-2">
              Payment receipt <span className="text-red-700">*</span>
            </Label>
            <div className="flex flex-col items-start space-y-4">
              <FileInput
                accept="image/png, image/jpeg"
                color="gray"
                sizing="lg"
                onChange={onSelectReceipt}
              />
              {receipt && (
                <div className="rounded-xl overflow-hidden h-40 aspect-[4/3] object-cover">
                  <img
                    src={preview}
                    className="w-full h-full m-auto object-cover"
                  />
                </div>
              )}
            </div>
          </div>

          <div>
            <Label className="block text-sm font-bold mb-2">
              Photo or Video proof <span className="text-red-700">*</span>
            </Label>
            <div className="flex flex-col items-start space-y-4">
              <FileInput
                accept="image/png, image/jpeg"
                color="gray"
                sizing="lg"
                onChange={onSelectMedia}
              />
            </div>
          </div>
        </div>

        <Label className="block text-sm font-bold mb-2">Description <span className="text-red-700">*</span></Label>
        <Textarea
          {...register("description", {
            required: "Description is required",
          })}
          rows={4}
          placeholder="Provide details of the fund spending and explain the uploaded proof"
        />
        {errors.description?.type === "required" && (
          <p className="text-red-600 text-sm">{errors.description.message}</p>
        )}

        <div className="flex space-x-1 mt-10">
          <Button type="submit" color="success" isProcessing={isLoading}>
            Upload Proof
          </Button>
          <Button color="light" onClick={() => setWriteOpen(false)}>
            Cancel
          </Button>
        </div>
      </form>

      <div className="my-5 space-y-4">
        <div className="font-semibold">
          Your spending proofs (
          {proofedMilestones ? proofedMilestones?.length : 0})
        </div>
        <div className="space-y-3">
          {proofedMilestones?.map((p) => (
            <div key={p.id}>
              <Card>
                <div className="flex justify-between">
                  <div className="flex gap-1">
                    <div>Milestone: </div>
                    <div className="font-bold">{p.milestone_title}</div>
                  </div>
                  <div>
                    <Tooltip content={ProofStatusMap[p.status].tooltip}>
                      <Badge color={ProofStatusMap[p.status].color}>
                        {p.status}
                      </Badge>
                    </Tooltip>
                  </div>
                </div>

                {p.status === "rejected" && (
                  <div>
                    Rejected for:{" "}
                    <span className="text-red-800">{`${p?.rejected_cause}`}</span>
                  </div>
                )}

                <div className="grid grid-cols-3 px-5 gap-5">
                  <div className="col-span-1">
                    <div>Transfer receipt: </div>
                    <div className="rounded-xl overflow-hidden aspect-[3/5] object-cover my-4">
                      <img
                        src={p.transfer_image}
                        className="w-full h-full m-auto object-cover"
                      />
                    </div>
                  </div>
                  <div className="col-span-2">
                    <div>Media proof: </div>
                    <div className="flex h-full justify-center items-center">
                      <div className="rounded-xl overflow-hidden aspect-[16/9] object-cover my-4">
                        <img
                          src={p.proof_media}
                          className="w-full h-full m-auto object-cover"
                        />
                      </div>
                    </div>
                  </div>
                </div>
                <div>
                  <span className="text-gray-600">Note: </span> {p.description}
                </div>
              </Card>
            </div>
          ))}
        </div>
      </div>

      <Modal
        show={isSuccessful}
        size="md"
        onClose={() => setIsSuccessful(false)}
        popup
      >
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img
              src="/success.svg"
              height={32}
              width={32}
              className="mx-auto"
            />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              New update posted!
            </h3>
          </div>
        </Modal.Body>
      </Modal>
      <Modal
        show={isFailed.status}
        size="md"
        onClose={() => setIsFailed(false)}
        popup
      >
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img src="/fail.svg" height={32} width={32} className="mx-auto" />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              Failed to post new update. Please try again later
            </h3>
          </div>
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default ManageProofs;
