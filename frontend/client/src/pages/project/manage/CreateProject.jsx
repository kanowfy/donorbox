import {
  Button,
  ButtonGroup,
  Checkbox,
  Datepicker,
  FileInput,
  Modal,
} from "flowbite-react";
import { CategoryIndexMap } from "../../../constants";
import { useEffect, useState } from "react";
import { useForm, useFieldArray } from "react-hook-form";
import Cleave from "cleave.js/react";
import utils from "../../../utils/utils";
import uploadService from "../../../services/upload";
import projectService from "../../../services/project";
import { useAuthContext } from "../../../context/AuthContext";
import { useNavigate } from "react-router-dom";
import MDEditor from "@uiw/react-md-editor";
import MilestoneInput from "../../../components/MilestoneInput";

const CreateProject = () => {
  const { token } = useAuthContext();
  const navigate = useNavigate();
  const [tosChecked, setTosChecked] = useState(false);
  const [img, setImg] = useState();
  const [preview, setPreview] = useState();
  const [geodata, setGeodata] = useState();
  const [country, setCountry] = useState();
  const [cityList, setCityList] = useState();
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);
  const [failedReason, setFailedReason] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [desc, setDesc] = useState("");
  const [step, setStep] = useState(1);

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    control,
  } = useForm();

  const { fields, append, remove } = useFieldArray({
    control,
    name: "milestones", // Name of the field array in the form
  });

  useEffect(() => {
    const loadFile = async () => {
      try {
        const res = await fetch("/countries-states.min.json");
        const data = await res.json();
        setGeodata(data);
      } catch (err) {
        console.error(err);
      }
    };

    if (!token) {
      navigate("/login?redirect=start-fundraiser");
    }

    loadFile();
  }, [navigate, token]);

  useEffect(() => {
    if (!img) {
      setPreview(undefined);
      return;
    }

    const objectUrl = URL.createObjectURL(img);
    setPreview(objectUrl);

    return () => URL.revokeObjectURL(img);
  }, [img]);

  function onSelectImage(e) {
    if (!e.target.files || e.target.files.length == 0) {
      setImg(undefined);
      return;
    }

    setImg(e.target.files[0]);
    setValue("img", e.target.files[0]);
  }

  const onSubmit = async (data) => {
    setIsLoading(true);
    try {
      const imageUrl = await uploadImage(data.img);
      const payload = {
        category_id: Number(data.category),
        title: data.title,
        description: desc,
        cover_picture: imageUrl,
        fund_goal: data.fund_goal,
        receiver_name: data.receiver_name,
        receiver_number: data.receiver_number,
        address: data.address,
        district: data.city,
        city: data.city,
        country: data.country,
        end_date: data.end_date,
        milestones: data.milestones,
      };
      console.log("payload: ", payload);
      const response = await projectService.create(token, payload);
      setIsLoading(false);
      setIsSuccessful(true);
      setTimeout(() => {
        setIsSuccessful(false);
        navigate(`/manage/${response.result.project.id}`);
      }, 3000);
    } catch (err) {
      //modal
      setFailedReason(err.response.data.error);
      setIsFailed(true);
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

  const handleSelectCountry = (e) => {
    setCountry(e.target.value);
    setCityList(utils.getCitiesByCountry(geodata, e.target.value));
  };

  return (
    <section className="py-10 flex flex-col items-center bg-gray-50">
      <div className="w-1/2 border rounded-lg shadow-xl px-16 py-5 bg-white mb-20">
        <div className="flex flex-col items-center">
          <div className="text-4xl font-bold mt-5 underline px-4 py-3 text-teal-500">
            Start a new Fundraiser
          </div>
          <div className="text-sm font-normal no-underline text-gray-600">
            Complete the form below to start a new fundraiser
          </div>
        </div>
        {/*{step === 1 && (
          <div className="flex justify-end">
            <Button
              onClick={() => setStep(2)}
              color={step === 2 ? "dark" : "light"}
            >
              Milestone Details
            <IoArrowForward className="ml-3 mt-1 h-4 w-4"/>
            </Button>
          </div>
        )}
        {step === 2 && (
          <div className="flex justify-start">
            <Button
              onClick={() => setStep(1)}
              color={step === 1 ? "dark" : "light"}
            >
            <IoArrowBack className="mr-3 mt-1 h-4 w-4"/>
              Basic Information
            </Button>
          </div>
        )}*/}
        <div className="flex justify-end mt-5 underline">
          {step === 1 ? "1/2 Basic Details" : "2/2 Milestone Details"}
        </div>

        <form
          className="space-y-4 md:space-y-6 my-10"
          onSubmit={handleSubmit(onSubmit)}
        >
          {step === 1 && (
            <>
              <div className="flex items-end space-x-3">
                <label className="block mb-2 font-medium text-gray-900">
                  Select category:{" "}
                </label>
                <select
                  {...register("category", {
                    required: "Category is required",
                  })}
                  name="category"
                  defaultValue=""
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5"
                >
                  <option value="" disabled></option>
                  {Object.entries(CategoryIndexMap).map(([name, num]) => (
                    <option value={num} key={num}>
                      {name}
                    </option>
                  ))}
                </select>
              </div>
              <div className="flex items-end space-x-3">
                <label className="block mb-2 font-medium text-gray-900">
                  Choose cover picture:
                </label>
                <FileInput
                  accept="image/png, image/jpeg"
                  onChange={onSelectImage}
                />
              </div>
              {img && (
                <div className="rounded-xl overflow-hidden h-40 aspect-[4/3] object-cover">
                  <img
                    src={preview}
                    className="w-full h-full m-auto object-cover"
                  />
                </div>
              )}
              <div>
                <label className="block mb-2 font-medium text-gray-900">
                  Title:
                </label>
                <input
                  {...register("title", {
                    required: "Title is required",
                  })}
                  type="text"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                  placeholder="Enter title"
                />
                {errors.title?.type === "required" && (
                  <p className="text-red-600 text-sm">{errors.title.message}</p>
                )}
              </div>
              <div data-color-mode="light">
                <label className="block mb-2 font-medium text-gray-900">
                  Description:
                </label>
                {/*<textarea
              {...register("description", {
                required: "Description is required",
              })}
              rows={20}
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
            />
            */}
                <MDEditor
                  value={desc}
                  onChange={(val) => setDesc(val || "")}
                  height={400}
                  required
                />
                {errors.description?.type === "required" && (
                  <p className="text-red-600 text-sm">
                    {errors.description.message}
                  </p>
                )}
              </div>
              <div>
                <label className="block mb-2 font-medium text-gray-900">
                  Set due date:
                </label>
                <Datepicker
                  minDate={new Date(Date.now())}
                  showClearButton={false}
                  showTodayButton={false}
                  autoHide={true}
                  onSelectedDateChanged={(date) =>
                    setValue("end_date", utils.getRFC3339DateString(date))
                  }
                  className="w-fit"
                />
              </div>

              <div className="p-5 border rounded-lg space-y-4">
                <div className="font-semibold">Beneficiary Information</div>
                <div className="flex space-x-4">
                  <div className="flex items-end space-x-3">
                    <label className="block mb-2 font-medium text-gray-900">
                      Name:{" "}
                    </label>
                    <input
                      {...register("receiver_name", {
                        required: "Beneficiary name is required",
                      })}
                      type="text"
                      className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                      placeholder="Enter name"
                      required
                    />
                    {errors.receiver_name?.type === "required" && (
                      <p className="text-red-600 text-sm">
                        {errors.receiver_name.message}
                      </p>
                    )}
                  </div>
                  <div className="flex items-end space-x-3">
                    <label className="block font-medium text-gray-900">
                      Phone number:{" "}
                    </label>
                    <input
                      {...register("receiver_number", {
                        required: "Beneficiary number is required",
                      })}
                      type="text"
                      className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                      placeholder="Enter number"
                      required
                    />
                    {errors.receiver_number?.type === "required" && (
                      <p className="text-red-600 text-sm">
                        {errors.receiver_number.message}
                      </p>
                    )}
                  </div>
                </div>

                <div className="flex items-end space-x-3">
                  <label className="block mb-2 font-medium text-gray-900">
                    Select location:{" "}
                  </label>
                  <select
                    {...register("country", {
                      required: "Country is required",
                    })}
                    name="country"
                    defaultValue=""
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5 max-w-2xl"
                    onChange={handleSelectCountry}
                  >
                    <option value="" disabled>
                      Choose country
                    </option>
                    {geodata &&
                      geodata.map((c, i) => (
                        <option value={c.name} key={i}>
                          {c.name}
                        </option>
                      ))}
                  </select>
                  <select
                    {...register("city", {
                      required: "city is required",
                    })}
                    name="city"
                    defaultValue=""
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5"
                    disabled={!country}
                  >
                    <option value="" disabled>
                      Choose city
                    </option>
                    {cityList?.map((p, i) => (
                      <option value={p} key={i}>
                        {p}
                      </option>
                    ))}
                  </select>
                </div>
                {errors.country?.type === "required" && (
                  <p className="text-red-600 text-sm">
                    {errors.country.message}
                  </p>
                )}
                {errors.city?.type === "required" && (
                  <p className="text-red-600 text-sm">{errors.city.message}</p>
                )}
                <div>
                  <input
                    {...register("address", {
                      required: "Address is required",
                    })}
                    type="text"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-2/3 p-2.5"
                    placeholder="Enter Address"
                    required
                  />
                  {errors.address?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errors.address.message}
                    </p>
                  )}
                </div>
              </div>
              {/*<div className="flex items-baseline space-x-1">
            <label className="block mb-2 font-medium text-gray-900">
              Set goal:
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
                {...register("goal_amount", {
                  required: "Donation amount is required",
                  min: 50000,
                  max: 100000000,
                })}
                onChange={(e) =>
                  setValue("goal_amount", e.target.value.replace(/,/g, ""))
                }
                className="border-0 focus:ring-0 focus:border-0 bg-gray-50 autofill:bg-gray-50"
                placeholder=""
              />
            </div>
            {errors.goal_amount?.type === "required" && (
              <p className="text-red-600 text-sm">
                {errors.goal_amount.message}
              </p>
            )}
          </div>*/}
              <div className="flex justify-end">
                <Button color="light" onClick={() => setStep(2)} size="xl">
                  Next
                </Button>
              </div>
            </>
          )}
          {step === 2 && (
            <div className="flex flex-col items-center space-y-5">
              <div className="w-full flex flex-col items-center space-y-5">
                {fields.map((field, index) => (
                  <MilestoneInput
                    index={index}
                    register={register}
                    setValue={setValue}
                    key={field.id}
                  />
                ))}
              </div>
              <ButtonGroup className="mb-10">
                <Button
                  color="failure"
                  pill
                  onClick={() => remove(fields.length - 1)}
                  disabled={fields.length <= 1}
                >
                  - Remove one
                </Button>
                <Button
                  color="light"
                  pill
                  onClick={() =>
                    append({
                      title: "",
                      description: "",
                      fund_goal: 0,
                      bank_description: "",
                    })
                  }
                >
                  + Add Milestone
                </Button>
              </ButtonGroup>

              <div className="flex text-sm items-end space-x-1">
                <Checkbox
                  className="h-5 w-5 mr-1 checked:bg-blue-700 focus:ring-0"
                  checked={tosChecked}
                  onChange={() => setTosChecked(!tosChecked)}
                />
                <div>
                  By starting a new fundraiser, you must agree with our platform
                  fund regulation
                </div>
                <div className="text-sm font-semibold underline">
                  Terms and Conditions
                </div>
              </div>
              <div className="w-full flex gap-2 justify-center">
                <Button color="light" onClick={() => setStep(1)} size="xl">
                  Back
                </Button>
                <Button
                  color="blue"
                  size="xl"
                  type="submit"
                  disabled={!tosChecked}
                  isProcessing={isLoading}
                >
                  Create
                </Button>
              </div>
            </div>
          )}
        </form>
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
            <h3 className="mb-5 text-xl font-normal text-gray-500 dark:text-gray-400">
              Fundraiser created successfully!
            </h3>
            <p className="text-xs text-gray-600">
              Redirecting to the fundraiser management page...{" "}
            </p>
          </div>
        </Modal.Body>
      </Modal>
      <Modal show={isFailed} size="md" onClose={() => setIsFailed(false)} popup>
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img src="/fail.svg" height={32} width={32} className="mx-auto" />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              Failed to create fundraiser:
            </h3>
            <h3 className="mb-5 text-red-500">{failedReason}</h3>
          </div>
        </Modal.Body>
      </Modal>
    </section>
  );
};

export default CreateProject;
