import { Button } from "flowbite-react";
import { useState } from "react";
import ProjectListCard from "../../../components/ProjectListCard";
import utils from "../../../utils/utils";

const ProjectList = () => {
  const [hasProject] = useState(false);
  return (
    <section className="py-10 flex flex-col items-center">
      <div className="flex justify-between w-2/5">
        <div className="text-3xl font-semibold my-2">Your Fundraisers</div>
        <Button color="green" className="h-fit" size="lg">
          Start a Fundraiser
        </Button>
      </div>
      <div className="w-2/5">
        {!hasProject ? (
          <div className="mt-20">
            <ProjectListCard
              id="69"
              title="my son go to school"
              img="https://images.gofundme.com/6KzAXSWgpeMlN5z4jC9UH1sda7k=/720x405/https://d2g8igdw686xgo.cloudfront.net/79205119_171177018484968_r.png"
              date={utils.formatDate(new Date(Date.now()))}
            />
          </div>
        ) : (
          <div className="flex justify-center text-lg mt-28">
            You have not started any fundraiser.
          </div>
        )}
      </div>
    </section>
  );
};

export default ProjectList;
