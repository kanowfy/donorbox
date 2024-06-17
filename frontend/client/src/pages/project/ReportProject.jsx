import { useEffect, useState } from "react";
import { Link, useNavigate, useParams } from "react-router-dom";
import projectService from "../../services/project";
import { Button } from "flowbite-react";
import { useForm } from "react-hook-form";
import Cleave from "cleave.js/react";
import "cleave.js/dist/addons/cleave-phone.i18n.js";

const reasons = [
  "Fundraiser contains factually incorrect information",
  "Fundraiser impersonates someone else or copies another fundraiser",
  "Fundraiser violates law",
  "Fundraiser contains copyrighted material",
];

const ReportProject = () => {
  const params = useParams();
  const [countries, setCountries] = useState();
  const [phoneCode, setPhoneCode] = useState("VN");
  const [project, setProject] = useState();
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors },
    setValue,
    reset,
  } = useForm();

  const onSubmit = (data) => {
    console.log(data);
    reset();
  };

  const handleSelectCountry = (e) => {
    const country = countries?.filter((c) => c.name == e.target.value)[0];
    setPhoneCode(country?.iso2);
  };

  useEffect(() => {
    const loadFile = async () => {
      try {
        const res = await fetch("/countries-states.min.json");
        const data = await res.json();
        setCountries(data);
      } catch (err) {
        console.error(err);
      }
    };

    const fetchProject = async () => {
      try {
        const response = await projectService.getOne(params.id);
        setProject(response.project);
      } catch (err) {
        navigate("/not-found");
        console.error(err);
      }
    };
    loadFile();
    fetchProject();
  }, [params.id, navigate]);

  return (
    <div className="w-full flex justify-center">
      <div className="w-1/3 my-10 space-y-7">
        <div className="text-4xl tracking-tight font-bold">
          Report a fundraiser
        </div>
        <form className="space-y-5" onSubmit={handleSubmit(onSubmit)}>
          <div className="space-y-3">
            <div>
              Provide us your contact information should we require more
              details.
            </div>
            <div>
              <label className="block  font-medium text-gray-600">
                Your full name:
              </label>
              <input
                {...register("full_name", {
                  required: "Your full name is required",
                })}
                type="text"
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                placeholder=""
              />
              {errors.full_name?.type === "required" && (
                <p className="text-red-600 text-sm">
                  {errors.full_name.message}
                </p>
              )}
            </div>
            <div>
              <label className="block  font-medium text-gray-600">
                Your Email:
              </label>
              <input
                {...register("email", {
                  required: "Your Email is required",
                })}
                type="email"
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                placeholder=""
              />
              {errors.email?.type === "required" && (
                <p className="text-red-600 text-sm">{errors.email.message}</p>
              )}
            </div>
            <div>
              <label className="block  font-medium text-gray-600">
                Your phone number:
              </label>
              <div className="grid grid-cols-10 gap-1">
                <select
                  name="country"
                  defaultValue=""
                  className="col-span-3 bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5"
                  onChange={handleSelectCountry}
                >
                  <option value="Vietnam">Vietnam</option>
                  {countries &&
                    countries.map((c, i) => (
                      <option value={c.name} key={i}>
                        {c.name}
                      </option>
                    ))}
                </select>
                <Cleave
                  options={{
                    phone: true,
                    phoneRegionCode: phoneCode,
                  }}
                  {...register("phone_number", {
                    required: "Your phone number is required",
                  })}
                  onChange={(e) =>
                    setValue("phone_number", e.target.value.replace(/\s/g, ""))
                  }
                  type="text"
                  className="col-span-7 bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                  placeholder=""
                />
              </div>
              {errors.phone_number?.type === "required" && (
                <p className="text-red-600 text-sm">
                  {errors.phone_number.message}
                </p>
              )}
            </div>
          </div>
          <hr />
          <div className="space-y-3">
            <div className="flex space-x-1">
              <div>Tell us why you are reporting </div>
              <Link
                to={`/fundraiser/${params.id}`}
                className="hover:underline font-bold text-gray-700"
              >
                {project?.title}
              </Link>
            </div>
            <div>
              <label className="block  font-medium text-gray-600">
                Reason:
              </label>
              <select
                {...register("reason", {
                  required: "Please provide a reason for your report",
                })}
                name="reason"
                defaultValue=""
                className="w-full bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5"
              >
                <option value="" disabled>
                  Please choose one
                </option>
                {reasons.map((r, i) => (
                  <option key={i} value={r}>
                    {r}
                  </option>
                ))}
              </select>
              {errors.reason?.type === "required" && (
                <p className="text-red-600 text-sm">{errors.reason.message}</p>
              )}
            </div>
            <div>
              <label className="block  font-medium text-gray-600">
                Details of the report:
              </label>
              <textarea
                {...register("details", {
                  required: "Please provide details for your report",
                })}
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                placeholder=""
                rows={4}
              />
              {errors.details?.type === "required" && (
                <p className="text-red-600 text-sm">{errors.details.message}</p>
              )}
            </div>
          </div>
          <Button color="dark" size="lg" type="submit">
            Submit report
          </Button>
        </form>
      </div>
    </div>
  );
};

export default ReportProject;
