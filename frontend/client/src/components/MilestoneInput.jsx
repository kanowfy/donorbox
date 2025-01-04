import Cleave from "cleave.js/react";

const MilestoneInput = ({ index, register, setValue, errors }) => {
  return (
    <div className="w-2/3 border p-3 rounded-lg">
      <div className="text-xl underline mb-2 font-semibold">
        Milestone {index + 1}:
      </div>
      {/* Title */}
      <div>
        <label className="block mb-2 font-medium text-gray-900">Title <span className="text-red-700">*</span></label>
        <input
          {...register(`milestones.${index}.title`, {
            required: "Title is required",
          })}
          type="text"
          className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
          placeholder="Enter title"
          required
        />
      </div>

      <div className="mt-2">
        <label className="block mb-2 font-medium text-gray-900">
          Details of the milestone
        </label>
        <textarea
          {...register(`milestones.${index}.description`)}
          rows={3}
          placeholder="Short description of the milestone"
          className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
        />
      </div>
      {/* Goal */}
      <div className="flex items-baseline space-x-1 mt-5">
        <label className="block mb-2 font-medium text-gray-900">
          Fund goal <span className="text-red-700">*</span>
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
            {...register(`milestones.${index}.fund_goal`, {
              required: "Donation amount is required",
              min: 50000,
              max: 100000000,
            })}
            onChange={(e) =>
              setValue(
                `milestones.${index}.fund_goal`,
                e.target.value.replace(/,/g, "")
              )
            }
            className="border-0 focus:ring-0 focus:border-0 bg-gray-50 autofill:bg-gray-50"
            placeholder=""
          />
        </div>
        {/*errors.goal_amount?.type === "required" && (
          <p className="text-red-600 text-sm">{errors.goal_amount.message}</p>
        )*/}
      </div>
      <div className="mt-2">
        <label className="block mb-2 font-medium text-gray-900">
          Bank description <span className="text-red-700">*</span>
        </label>
        <textarea
          {...register(`milestones.${index}.bank_description`, {
            required: "Bank description is required",
          })}
          rows={3}
          placeholder="Please specify how we can send donated fund to the beneficiary"
          className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
        />
      </div>
    </div>
  );
};

export default MilestoneInput;
