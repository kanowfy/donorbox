import { Button, Datepicker, FileInput } from "flowbite-react";
import { CategoryIndexMap } from "../../../constants";
import { useEffect, useState } from "react";

const CreateProject = () => {
  const [img, setImg] = useState();
  const [preview, setPreview] = useState();

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
  }

  return (
    <section className="py-10 flex flex-col items-center">
      <div className="flex justify-center">
        <div className="text-4xl font-semibold tracking-wide my-2">
          Start a new Fundraiser
        </div>
      </div>
      <div className="w-1/2 border rounded-lg shadow-xl px-16 py-5">
        <form className="space-y-4 md:space-y-6 my-10">
          <div className="flex items-end space-x-3">
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Select category:{" "}
            </label>
            <select
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
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Choose cover picture:
            </label>
            <FileInput
              accept="image/png, image/jpeg"
              onChange={onSelectImage}
            />
          </div>
          {img && (
            <div className="h-60 aspect-[4/3]">
              <img src={preview} />
            </div>
          )}
          <div>
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Title:
            </label>
            <input
              type="text"
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
              placeholder="Enter title"
            />
          </div>
          <div>
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Description:
            </label>
            <textarea
              rows={20}
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
            />
          </div>
          <div className="flex items-end space-x-3">
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Select location:{" "}
            </label>
            <select
              name="category"
              defaultValue=""
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5"
            >
              <option value="" disabled>
                Choose country
              </option>
              {Object.entries(CategoryIndexMap).map(([name, num]) => (
                <option value={num} key={num}>
                  {name}
                </option>
              ))}
            </select>
            <select
              name="category"
              defaultValue=""
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block p-2.5"
            >
              <option value="" disabled>
                Choose city
              </option>
              {Object.entries(CategoryIndexMap).map(([name, num]) => (
                <option value={num} key={num}>
                  {name}
                </option>
              ))}
            </select>
          </div>
          <div className="flex items-end space-x-3">
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Set goal:
            </label>
            <input
              type="text"
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-1/3 p-2.5"
            />
            <span className="text-gray-800 text-sm">₫</span>
          </div>
          <div>
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Set due date:
            </label>
            <Datepicker
              minDate={new Date(Date.now())}
              showClearButton={false}
              showTodayButton={false}
              inline
            />
          </div>
          <div className="flex justify-center pt-10">
            <Button color="teal" className="w-1/3" size={"xl"} type="submit">
              Create
            </Button>
          </div>
        </form>
      </div>
    </section>
  );
};

export default CreateProject;
