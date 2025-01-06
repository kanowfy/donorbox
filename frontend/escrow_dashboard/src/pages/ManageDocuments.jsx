import { useForm, useFieldArray } from "react-hook-form";
import { Button, Modal } from "flowbite-react";
import ragService from "../services/rag";
import { useAuthContext } from "../context/AuthContext";
import { useState } from "react";
import { useNavigate } from "react-router-dom";

const ManageDocuments = () => {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);
  const { token } = useAuthContext();
  const {
    register,
    handleSubmit,
    formState: { errors },
    control,
  } = useForm({
    defaultValues: {
      documents: [{ text: "" }],
    },
  });

  const { fields, append, remove } = useFieldArray({
    control,
    name: "documents",
  });

  const onSubmit = async (data) => {
    console.log("data ", data);
    setIsLoading(true);
    try {
      await ragService.addDocument(token, data);
      setIsLoading(false);
      setIsSuccessful(true);
      setTimeout(() => {
        navigate(0);
      }, 5000);
    } catch (err) {
      setIsLoading(false);
      setIsFailed(true);
      console.error(err);
    }
  };

  return (
    <div className="p-10 bg-slate-200 w-full space-y-10 font-sans min-h-screen">
      <div className="text-3xl font-semibold tracking-tight">
        Add Documents to Knowledge Base
      </div>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="bg-gray-100 border rounded-lg p-10"
      >
        <div>
          <div className="w-full space-y-5">
            <label className="block text-lg font-medium text-gray-900">
              Document:
            </label>
            {fields.map((field, index) => (
              <div key={field.id}>
                <textarea
                  {...register(`documents.${index}.text`)}
                  rows={3}
                  placeholder=""
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                />
              </div>
            ))}
          </div>
          <div className="flex justify-center gap-2 my-4">
            <Button
              size="sm"
              color="light"
              className="font-bold"
              onClick={() => remove(fields.length - 1)}
              disabled={fields.length <= 1}
            >
              -
            </Button>
            <Button
              size="sm"
              color="dark"
              className="font-bold"
              onClick={() =>
                append({
                  text: "",
                })
              }
            >
              +
            </Button>
          </div>
        </div>
        <div className="flex justify-center">
          <Button type="Submit" color="blue" size="xl" isProcessing={isLoading}>
            Add Documents
          </Button>
        </div>
      </form>
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
              Action completed
            </h3>
          </div>
        </Modal.Body>
      </Modal>
      <Modal show={isFailed} size="md" onClose={() => setIsFailed(false)} popup>
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img src="/fail.svg" height={32} width={32} className="mx-auto" />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              Failed to complete action
            </h3>
          </div>
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default ManageDocuments;
