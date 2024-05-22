import { Button, FileInput, Textarea } from "flowbite-react";
import { useEffect, useState } from "react";
import ProjectUpdate from "../../../components/ProjectUpdate";

const ManageUpdates = () => {
  const [photo, setPhoto] = useState();
  const [preview, setPreview] = useState();

  const [writeOpen, setWriteOpen] = useState(false);
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
      <form className={`my-5 space-y-3 ${writeOpen ? "" : "hidden"}`}>
        <Textarea rows={4} placeholder="Write an update" />
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
        <Button type="submit" color="dark">
          Post update
        </Button>
      </form>

      <div className="my-5 space-y-4">
        <div className="font-medium">Your updates ({0})</div>
        <div className="space-y-3">
          <ProjectUpdate
            first_name="Dung"
            last_name="Nguyen"
            content="Lorem ipsum dolor sit amet consectetur adipisicing elit. Quos eos impedit quisquam dolorem corrupti consectetur, rerum cupiditate laboriosam? Molestias alias quia doloremque voluptates ipsum molestiae ducimus cupiditate nisi natus quaerat."
            created_at="2024-05-05T08:51:04+00:00"
          />

          <ProjectUpdate
            first_name="Dung"
            last_name="Nguyen"
            content="Lorem ipsum dolor sit amet consectetur adipisicing elit. Quos eos impedit quisquam dolorem corrupti consectetur, rerum cupiditate laboriosam? Molestias alias quia doloremque voluptates ipsum molestiae ducimus cupiditate nisi natus quaerat."
            created_at="2024-05-05T08:51:04+00:00"
          />
        </div>
      </div>
    </div>
  );
};

export default ManageUpdates;
