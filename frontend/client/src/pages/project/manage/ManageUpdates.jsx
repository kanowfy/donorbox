import { Button, FileInput, Modal, Textarea } from "flowbite-react";
import { useEffect, useState } from "react";
import ProjectUpdate from "../../../components/ProjectUpdate";
import { useForm } from "react-hook-form";
import uploadService from "../../../services/upload";
import projectService from "../../../services/project";
import { useNavigate, useOutletContext } from "react-router-dom";
import { useAuthContext } from "../../../context/AuthContext";

const ManageUpdates = () => {
  const { token, user } = useAuthContext();
  const { project } = useOutletContext();
  const navigate = useNavigate();
  const [photo, setPhoto] = useState();
  const [preview, setPreview] = useState();
  const [fundraiserUpdates, setFundraiserUpdates] = useState();
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
    const fetchUpdates = async () => {
      try {
        const response = await projectService.getUpdates(project.id);
        setFundraiserUpdates(response.updates);
      } catch (err) {
        console.error(err);
      }
    };

    fetchUpdates();
  }, [project]);

  useEffect(() => {
    if (!photo) {
      setPreview(undefined);
      return;
    }

    const objectUrl = URL.createObjectURL(photo);
    setPreview(objectUrl);

    return () => URL.revokeObjectURL(photo);
  }, [photo]);

  function onSelectImage(e) {
    if (!e.target.files || e.target.files.length == 0) {
      setPhoto(undefined);
      return;
    }

    setPhoto(e.target.files[0]);
  }

  const onSubmit = async (data) => {
    setIsLoading(true);
    try {
      const payload = {
        project_id: project.id,
        description: data.description,
      };

      if (photo) {
        payload.attachment_photo = await uploadImage(photo);
      }

      await projectService.createUpdate(token, payload);
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
      <div className="text-3xl font-semibold">Updates</div>
      <div className="text-gray-600 text-sm mt-1">
        Keep your donors up to date!
      </div>
      <div
        className={`border rounded-lg hover:cursor-text p-4 text-sm my-5 text-gray-600 hover:text-gray-800 hover:bg-gray-100 hover:border-gray-800 ${
          writeOpen ? "hidden" : ""
        }`}
        onClick={() => setWriteOpen(true)}
      >
        Write an update
      </div>
      <form
        className={`my-5 space-y-3 ${writeOpen ? "" : "hidden"}`}
        onSubmit={handleSubmit(onSubmit)}
      >
        <Textarea
          {...register("description", {
            required: "Description is required",
          })}
          rows={4}
          placeholder="Write an update"
        />
        {errors.description?.type === "required" && (
          <p className="text-red-600 text-sm">{errors.description.message}</p>
        )}

        <label className="block text-sm font-medium">
          Add a photo (Optional)
        </label>
        <div className="flex flex-col items-start space-y-4">
          <FileInput
            accept="image/png, image/jpeg"
            color="gray"
            sizing="lg"
            onChange={onSelectImage}
          />
          {photo && (
            <div className="rounded-xl overflow-hidden h-40 aspect-[4/3] object-cover">
              <img
                src={preview}
                className="w-full h-full m-auto object-cover"
              />
            </div>
          )}
        </div>
        <div className="flex space-x-1">
          <Button type="submit" color="dark" isProcessing={isLoading}>
            Post update
          </Button>
          <Button color="light" onClick={() => setWriteOpen(false)}>
            Cancel
          </Button>
        </div>
      </form>

      <div className="my-5 space-y-4">
        <div className="font-medium">
          Your updates ({fundraiserUpdates ? fundraiserUpdates?.length : 0})
        </div>
        <div className="space-y-3">
          {fundraiserUpdates?.map((u) => (
            <ProjectUpdate
              key={u.id}
              first_name={user.first_name}
              last_name={user?.last_name}
              content={u.description}
              photo={u?.attachment_photo}
              created_at={u.created_at}
            />
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

export default ManageUpdates;
