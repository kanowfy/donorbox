import { Button, Checkbox, Textarea, Tooltip } from "flowbite-react";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import Cleave from "cleave.js/react";
import { useAuthContext } from "../../../context/AuthContext";
import { Link, useNavigate, useParams } from "react-router-dom";
import projectService from "../../../services/project";
import { FaAngleLeft } from "react-icons/fa6";
import { BsFillQuestionCircleFill } from "react-icons/bs";

const Donate = () => {
  const params = useParams();
  const navigate = useNavigate();
  const { user, loading } = useAuthContext() || {};
  const [wosChecked, setWosChecked] = useState(false);
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm();
  const [project, setProject] = useState();

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const response = await projectService.getOne(params.id);
        setProject(response.project);
      } catch (err) {
        navigate("/not-found");
        console.error(err);
      }
    };

    if (!loading) {
      fetchProject();
    }
  }, [params.id, loading, navigate]);

  const onSubmit = (data) => {
    console.log("data: ", data);
    navigate("/fundraiser/" + params.id + "/payment/checkout", {
      state: {
        amount: data.amount,
        project_id: params.id,
        user_id: user?.id,
        word_of_support: data?.word_of_support,
      },
    });
  };

  return (
    <section className="py-10 flex flex-col items-center bg-gray-100 min-h-screen">
      <div className="w-full grid grid-cols-3">
        <Link
          to={`/fundraiser/${params?.id}`}
          className="flex hover:underline justify-center"
        >
          <FaAngleLeft className="w-7 h-7" />
          <div>Back to Fundraiser</div>
        </Link>
        <Link
          to={`/fundraisers/${params?.id}`}
          className="flex items-center justify-center mb-6 text-2xl font-semibold text-gray-800 tracking-tight"
        >
          <img className="w-8 h-8 mr-2" src="/logo.svg" alt="logo" />
          Donorbox
        </Link>
      </div>
        <form
          className="min-w-1/3 mt-20 rounded-lg shadow-lg px-10 py-5 space-y-4 bg-white"
          onSubmit={handleSubmit(onSubmit)}
        >
          <div className="grid grid-cols-12 space-x-2">
            <div className="col-span-3 rounded-xl overflow-hidden h-20 aspect-[4/3] object-cover">
              <img
                src={project?.cover_picture}
                className="w-full h-full m-auto object-cover"
              />
            </div>
            <div className="col-span-9">
              <span>You are supporting </span>
              <span className="font-semibold">{project?.title}</span>
            </div>
          </div>
          <div>
            <label className="block mb-2 font-medium text-gray-900">
              Select amount to donate <span className="text-red-700">*</span>
            </label>
            <div className="flex items-baseline bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 w-full px-5 py-3">
              <span className="block mb-1 font-medium">₫</span>
              <Cleave
                options={{
                  numeral: true,
                  numericOnly: true,
                  numeralThousandsGroupStyle: "thousand",
                  numeralPositiveOnly: true,
                }}
                {...register("amount", {
                  required: "Donation amount is required",
                  min: 50000,
                  max: 50000000,
                })}
                onChange={(e) =>
                  setValue("amount", e.target.value.replace(/,/g, ""))
                }
                className="border-0 focus:ring-0 focus:border-0 bg-gray-50 w-11/12 text-xl font-medium autofill:bg-gray-50"
                placeholder=""
              />
            </div>
            {errors.amount?.type === "required" && (
              <p className="text-red-600 text-sm">{errors.amount.message}</p>
            )}
            {errors.amount?.type === "min" && (
              <p className="text-red-600 text-sm">
                Minimum donation amount is ₫50,000
              </p>
            )}
            {errors.amount?.type === "max" && (
              <p className="text-red-600 text-sm">
                Maximum donation amount is ₫50,000,000
              </p>
            )}
          </div>
        <div className="font-medium text-gray-700 flex space-x-1">
          <Checkbox
            onChange={() => setWosChecked(!wosChecked)}
            checked={wosChecked}
            className="mt-1 mr-1"
          />
          <div>Write word of support</div>
          <div className="mt-1">
            <Tooltip content="Your word of support will be displayed on the fundraiser page once you have made a successful donation.">
              <BsFillQuestionCircleFill className="w-4 h-4" />
            </Tooltip>
          </div>
        </div>

        {wosChecked && (
          <div>
            <Textarea
              {...register("word_of_support", {
                required: false,
                minLength: 2,
              })}
              placeholder="Enter word of support"
              rows={4}
            />
          </div>
        )}

          <Button
            color="success"
            className="w-full"
            size={"xl"}
            type="submit"
          >
            Proceed to Payment
          </Button>
        </form>
    </section>
  );
};

export default Donate;
