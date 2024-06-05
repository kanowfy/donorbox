import { Link } from "react-router-dom";
import ProjectCard from "../components/ProjectCard";
import { useEffect, useState } from "react";
import projectService from "../services/project";
import { FaSearch } from "react-icons/fa";

const Home = () => {
  const [projects, setProjects] = useState([]);

  useEffect(() => {
    const fetchProjects = async () => {
      try {
        const resp = await projectService.getAll();
        console.log(resp.projects);
        setProjects(resp.projects);
      } catch (err) {
        console.error(err);
      }
    };

    fetchProjects();
  }, []);

  return (
    <>
      <div>
        <section className="flex flex-col pt-20 items-center h-[33rem] bg-cover bg-gradient-to-b from-white to-sky-200">
          <div className="pb-7 font-semibold text-sky-700 text-6xl">
            Help those in need today
          </div>
          <div className="text-sky-800 pb-10 font-medium text-xl">
            Your home for communities, charities and people you care about
          </div>
          <div>
            <Link to="/start-fundraiser">
              <button className="mt-10 px-8 py-4 text-white font-medium rounded-xl shadow-lg bg-gradient-to-t from-cyan-500 to-blue-500 hover:bg-gradient-to-b">
                Start a Fundraiser
              </button>
            </Link>
          </div>
        </section>

        <section>
          <div className="min-h-screen px-10 pt-6">
            <div className="flex justify-between mx-48">
              <div className="font-medium text-2xl tracking-tight">
                Trending fundraisers
              </div>
              <div>
                <Link to="/search">
                  <div className="underline font-semibold text-xl text-gray-800 hover:text-sky-700 flex items-center space-x-2">
                    <div>Explore</div>
                    <FaSearch className="w-4 h-4 mt-1" />
                  </div>
                </Link>
              </div>
            </div>
            <div className="flex justify-center">
              <div className="grid grid-cols-1 gap-7 md:grid-cols-3 xl:grid-cols-4 mx-16 mt-10 mb-16">
                {projects.slice(0, 8).map((p) => (
                  <ProjectCard
                    id={p.id}
                    title={p.title}
                    cover={p.cover_picture}
                    currentAmount={p.current_amount}
                    goalAmount={p.goal_amount}
                    numBackings={p.backing_count}
                    key={p.id}
                  />
                ))}
              </div>
            </div>
          </div>
        </section>
      </div>
    </>
  );
};

export default Home;
