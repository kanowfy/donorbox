import { useParams } from "react-router-dom";
import projectService from "../../services/project";
import { useState, useEffect } from "react";
import ProjectCard from "../../components/ProjectCard";

const Category = () => {
  const params = useParams();
  const [category, setCategory] = useState();
  const [projects, setProjects] = useState();

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await projectService.getCategoryByName(params.name);
        setCategory(response.category);

        const projectResponse = await projectService.getByCategoryID(
          response.category.id
        );
        setProjects(projectResponse.projects.filter(p => ["ongoing", "finished"].includes(p.status)));

        console.log(projectResponse.projects);
      } catch (err) {
        console.error(err);
      }
    };

    fetchData();
  }, [params.name]);

  return (
    <div>
      <section className="flex justify-center mx-10 mt-14 mb-28">
        <div className="grid grid-cols-5 pt-16 gap-20">
          <div className="col-span-3 flex flex-col items-center justify-center">
            <div>
              <div className=" text-6xl font-bold">
                {`${
                  String(category?.name).charAt(0).toUpperCase() +
                  String(category?.name).slice(1)
                } Fundraising`}
              </div>
              <div className="text-gray-700 text-3xl">
                {category?.description}
              </div>
            </div>
          </div>
          <div className="col-span-2 rounded-xl max-w-96 overflow-hidden aspect-[4/5] object-cover border outline-1 shadow-xl">
            <img
              src={category?.cover_picture}
              className="w-full h-full m-auto object-cover"
            />
          </div>
        </div>
      </section>
      <hr/>
      <section className="flex flex-col items-center mx-36 mt-10 mb-20">
        <div className="text-2xl text-black font-bold">Explore {category?.name} fundraisers</div>
        <div className="flex justify-center">
          <div className="grid grid-cols-1 gap-10 md:grid-cols-2 xl:grid-cols-3 mx-16 mt-10 mb-16">
            {projects &&
              projects
                .slice(0, 8)
                .map((p) => (
                  <ProjectCard
                    id={p.id}
                    title={p.title}
                    cover={p.cover_picture}
                    totalFund={p.total_fund}
                    fundGoal={p.fund_goal}
                    numBackings={p.backing_count}
                    key={p.id}
                  />
                ))}
          </div>
        </div>
      </section>
    </div>
  );
};

export default Category;
