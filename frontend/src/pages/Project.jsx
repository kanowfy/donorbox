import { Avatar } from "flowbite-react";
import { Link } from "react-router-dom";
import Support from "../components/Support";
import DonateBox from "../components/DonateBox";

const Project = () => {
  return (
    <div className="mx-auto">
      <div className="text-3xl font-bold m-5">
        Callum and Jake Robinson, Support the Robinson Family
      </div>
      <div className="grid grid-cols-3 gap-4 ">
        <div className="col-span-2">
          <div className="rounded-xl overflow-hidden">
            <img
              src="https://w.wallhaven.cc/full/we/wallhaven-wer2x7.jpg"
              className="h-128"
            ></img>
          </div>
          <div className="m-3 flex justify-start">
            <Avatar
              alt="User settings"
              img="https://flowbite.com/docs/images/people/profile-picture-5.jpg"
              rounded
            >
              <div>
                <span className="flex justify-start tracking">
                  I am organizing this fundraiser.
                </span>
                <span className="block font-normal text-gray-500 text-sm">
                  Created 6d ago.
                </span>
              </div>
            </Avatar>
          </div>
          <div className="h-px bg-gray-300"></div>

          <p className="max-w-3xl tracking-tight mt-4">
            Lorem ipsum dolor sit amet, consectetur adipiscing elit. In sagittis
            ex aliquet velit facilisis pulvinar. Nam elementum fringilla tortor,
            a tincidunt nulla dapibus ut. Proin porttitor posuere ante, a
            fermentum tortor pulvinar in. Aenean quis lobortis diam, iaculis
            pulvinar mi. Vestibulum vehicula risus sed maximus gravida.
            Phasellus quis massa quam. Praesent vestibulum massa sed massa
            aliquam, finibus lacinia est egestas. Nam semper nisl non felis
            facilisis porttitor. Suspendisse potenti. Aenean dapibus lorem sit
            amet diam vehicula placerat. Mauris viverra dolor velit, a bibendum
            ipsum luctus ut. Fusce elementum, arcu ac ultrices scelerisque,
            sapien nulla lobortis nunc, et lobortis ante eros sit amet felis.
            Donec mollis ipsum sed turpis venenatis interdum. Vivamus porta orci
            laoreet mauris consequat, ut venenatis augue bibendum. Fusce eu
            lorem mattis ante commodo rutrum quis et magna. Curabitur erat
            turpis, feugiat eget massa volutpat, laoreet congue leo. Duis
            egestas tincidunt lorem, non tristique nunc lobortis eget. Nullam id
            viverra lectus. Vivamus pharetra consequat risus, sit amet ultrices
            sapien lacinia ac. Mauris id orci sed dolor porta rutrum sed eu
            eros. Maecenas neque erat, rutrum lacinia semper eget, dapibus et
            velit. Nullam elementum purus non nulla suscipit porttitor. Ut nunc
            massa, tristique ac quam id, mattis aliquam mauris. Quisque viverra
            a lacus id maximus. Donec vel malesuada elit. Nullam erat mi,
            ultricies at metus at, dapibus malesuada turpis. Curabitur sapien
            orci, pellentesque faucibus laoreet eget, porttitor quis enim. Fusce
            varius, ex ut viverra hendrerit, lorem elit vehicula dui, a gravida
            lacus lacus sed tortor. Nunc iaculis lectus ultrices sem porttitor,
            vitae posuere est scelerisque. Sed eget condimentum dui, a ultrices
            neque. Nulla eleifend lorem vel nisl dapibus, eu aliquam velit
            mattis. Donec at ullamcorper magna. Quisque ac vulputate libero.
            Aliquam tincidunt dapibus faucibus. Sed aliquet odio at mi feugiat,
            eu egestas purus interdum. Vestibulum molestie metus neque, sed
            vestibulum tortor feugiat sit amet. Vivamus sagittis ut magna non
            scelerisque. Sed sed euismod est, sit amet tincidunt urna. Praesent
            molestie ipsum lorem, a viverra nisl gravida sit amet. Curabitur
            luctus aliquet urna, id euismod risus ultrices vitae. Quisque
            aliquam lobortis diam, vitae dapibus tortor gravida et. Cras
            pulvinar cursus porttitor. Praesent aliquet eu massa ac imperdiet.
            Nunc pharetra purus vitae sapien convallis ultrices. Suspendisse eu
            lacinia mauris, eget eleifend dolor. Interdum et malesuada fames ac
            ante ipsum primis in faucibus. Quisque a nunc aliquet, tempus turpis
            vel, congue urna. Fusce porta bibendum quam id dignissim. Fusce eu
            purus tincidunt, rutrum quam quis, malesuada magna. Nam et ultrices
            velit. Vestibulum ante ipsum primis in faucibus orci luctus et
            ultrices posuere cubilia curae; Duis sit amet neque dui. Aenean
            aliquet elit in nulla finibus pretium. Donec maximus ullamcorper
            diam non efficitur.
          </p>
          <div className="flex justify-center my-5">
            <Link to="#">
              <div className="border text-xl flex py-3 max-w-lg rounded-lg border-gray-400 px-40 hover:bg-gray-100 hover:border-gray-900 duration-300">
                Donate
              </div>
            </Link>
          </div>

          <div className="h-px bg-gray-300"></div>

          <div className="my-5">
            <div className="text-xl font-semibold">
              Donators&apos; words of support
            </div>

            <Support
              avatar="https://flowbite.com/docs/images/people/profile-picture-5.jpg"
              amount={250000}
              day_since={21}
              comment="
              consectetur adipisicing elit. Velit quas explicabo hic possimus
              nisi placeat recusandae quo illum, fugit officia saepe, laudantium
              numquam rem quibusdam nulla nesciunt nobis reiciendis quos?"
            />
            <Support
              avatar="https://flowbite.com/docs/images/people/profile-picture-5.jpg"
              amount={250000}
              day_since={21}
              comment="skibidi toilet let go"
            />
            <Support
              avatar="https://flowbite.com/docs/images/people/profile-picture-5.jpg"
              amount={250000}
              day_since={21}
              comment="
              consectetur adipisicing elit. Velit quas explicabo hic possimus
              nisi placeat recusandae quo illum, fugit officia saepe, laudantium
              numquam rem quibusdam nulla nesciunt nobis reiciendis quos?"
            />
          </div>
        </div>
        <div className="col-span-1">
          <DonateBox />
        </div>
      </div>
    </div>
  );
};

export default Project;
