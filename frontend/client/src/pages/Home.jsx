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
        <section className="relative h-[48rem]">
          <div
            class="absolute inset-0 bg-top bg-cover bg-no-repeat filter brightness-50"
            style={{ backgroundImage: "url('/handshake.jpeg')" }}
          ></div>

          <div className="relative z-10 flex items-center flex-col top-36">
            <div className="pb-7 font-semibold text-green-100 text-7xl">
              Help those in need today
            </div>
            <div className="text-green-100 font-medium text-2xl">
              Your home for communities, charities and people you care about
            </div>

            {/*<div>
            <Link to="/start-fundraiser">
              <button className="text-lg mt-10 px-8 py-4 text-white font-semibold rounded-xl shadow-lg bg-gradient-to-t from-cyan-500 to-blue-500 hover:bg-gradient-to-b">
                Start a Fundraiser
              </button>
            </Link>
          </div>*/}
          </div>
        </section>

        <section>
          <div className="min-h-screen px-10 pt-6 bg-gray-50">
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
              <div className="grid grid-cols-1 gap-7 md:grid-cols-2 xl:grid-cols-4 mx-16 mt-10 mb-16">
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
          </div>
        </section>
      </div>
    </>
  );
};

export default Home;
