# [d0iasm.github.io](https://d0iasm.github.io/)

<img src="me.jpg" width=200vw height=200vw style="float:right;" />

**Asami Doi (土井麻未)**

Master's student researching on artificial life (ALife) in Japan. I'm interested in low-level programming especially an emulator, an OS, and a compiler.

Love programming ❤️

[Twitter](https://twitter.com/d0iasm) | [GitHub](https://github.com/d0iasm) | [Tech blog](https://d0iasm.github.io/blog/) | [note (jp)](https://note.com/d0iasm)


## Experience
**Internship at Wasmer** | *December 2019 - Now*
- Wasmer is the start-up company based on San Francisco, California. They are making the WebAssembly runtime. I Worked on the language integration project, one of the foreign function interface (FFI) technologies. My project enables us to execute WebAssembly functions from another language.

**Google Summer of Code 2019 at coreboot project** | *May 2019 - July 2019*
- Worked on the coreboot, one of the open source projects, which aims at replacing proprietary firmware (BIOS or UEFI). My project was "Adding QEMU/AArch64 Support to Coreboot". Coreboot already supported an ARMv8 archtecture but it didn't have a QEMU board support for ARMv8. I added the support for QEMU/AArch64 to coreboot by C language and assembly language.
  - [Project details](https://summerofcode.withgoogle.com/archive/2019/projects/4518359851335680/)
  - [Wrap-up blog](https://blogs.coreboot.org/blog/2019/08/23/gsoc-wrap-up-for-adding-qemu-aarch64-support-to-coreboot/)

**Internship at Google** | *August 2018 - November 2018*
- Worked on Chromium browser project in Google Japan. The projecrt was "ES Modules for Service Workers" that enable web developers to import JavaScript as a module system. I used C++ for the main implementation and JavaScript for unit tests and integrated tests.
  - [Chromium Gerrit](https://chromium-review.googlesource.com/q/owner:asamidoi)
  - [Design document](https://docs.google.com/document/d/1SeQ085YdBTtW3D_ygSpO0Wz2DAe8QiS1gj37IG5lstg/)
  - [Slide](https://docs.google.com/presentation/d/19mLH9FK5mOXlcQQkAb1QSeP7BZFV4ZRxNigX55EGeOA/)

**Backend Engineer at Tryeting** | *March 2018 - May 2019*
- Add many features to a web application in C#. Our service enables users to use AI online through APIs or website.

**Internship at pixiv** | *February 2018*
- Implemented the feature to pick up users who should receive notification or not in PHP.

**Frondend Engineer at Optimind** | *December 2018 - February 2018*
- Implemented the web application which shows us the optimized route to the destination on Google Maps by using React+Redux. This service is basically used by delivery people to save thier time.

**Internship at Google** | *August 2017 - October 2017*
- Created the web application in Python, Flask, anad Google App Engine to make it easy to analyze logs.

**Backend Engineer at PREVENT** | *April 2017 - July 2017*
- Created the web APIs in Ruby to transfer data between an iOS application and a web application by using Ruby on Rails, Heroku, and AWS.

**Internship at Ateam** | *September 2016*
- Created the prototype of a web service which we thought of from scratch in a group. Implemented it in Ruby by using Ruby on Rails.

**Director at university festival** | *April 2016 - October 2016*
- I was a leader of a group to create a website for a university festival in HTML and CSS. I mainly did project management and taught other students how to write code.
  - [Geikosai 2016](http://geikousai-ncu.com/2016/)

**Training Engineer at Kurusugawa Computer Inc.** | *August 2016*
- Trained as a software engineer for a month by creating Java libraries. However, I gave up the training because of lack of programming ability.

**Unity Mentor at Life is Tech** | *June 2016*
- Participated in Life is Tech as a Unity mentor. Made original games in C#.

**Member in UNITED** | *April 2015 - April 2016*
- This was the first opportunity to write code in HTML, CSS, JavaScript, and PHP. I was a member of UNITED that students organized for learning web by creating websites and web applications. We met once per a week, and sometimes we were commissioned for projects from other organizations.
  - [united.nagoya](http://united.nagoya/)

## Education
**Master of Informatics at Nagoya University** | *April 2018 - March 2020*
- Working on the research project in ALife field and studying about computer science. In my research, I'm trying to understand the dynamics of human's relationship based on self-driven particle system.

**Bachelor of Design at Nagoya City University** | *April 2014 - March 2018*
- Learned about UI/UX design, product design, and design thinking which is the process to understand users and redefine problems in order to produce valuable solutions.

## Research
- Asami Doi, Reiji Suzuki, and Takaya Arita, "A particle swarm model of socio-psychological dynamics based on Heider's balance theory”, Proceedings of the 25th International Symposium on Artificial Life and Robotics (AROB 2020), pp. 528-533. | *January 2020*
- Poster presentation with "[バランス理論に基づく非対称相互作用による社会的粒子群モデル](https://drive.google.com/file/d/1mGEK0KPO0QEur6excFx6EurVf4Ytef47/view?usp=sharing)" at 名古屋大学若手女性研究者サイエンスフォーラム | *August 2019*
- Poster presentation with "[Graphy: Chat Bot with Evolutionary Concept Graph using Dependency Analysis](http://www.interaction-ipsj.org/proceedings/2018/data/pdf/2P07.pdf)" at Interaction 2018 | *March 2018*

## Personal Projects
**[rvemu](https://github.com/d0iasm/rvemu)** | *October 2019 - Now*
- RISC-V online emulator with WebAssembly generated by Rust. You can execute RISC-V binaries on your browser.
- Related articles:
  - [Emulate 32-Bit And 64-Bit RISC-V In Your Browser With Asami’s Open Source rvemu | Gareth Halfacree, Hackster.io](https://riscv.org/2020/01/emulate-32-bit-and-64-bit-risc-v-in-your-browser-with-asamis-open-source-rvemu-gareth-halfacree-hackster-io/)
  - [Emulate 32-Bit and 64-Bit RISC-V in Your Browser with Asami's Open Source rvemu](https://www.hackster.io/news/emulate-32-bit-and-64-bit-risc-v-in-your-browser-with-asami-s-open-source-rvemu-b783f672e463)

**[x86emu](https://github.com/d0iasm/x86emu)** | *October 2019*
- A x86 emulator in Rust. This is a porting from the implementation written by C in "自作エミュレータで学ぶx86アーキテクチャ". It can now execute a binary file to calculate fibonacci numbers. There is a binary file for calculating a fibonacci number 10 (=55) in the test directory.

**[minigo](https://github.com/d0iasm/minigo)** | *August 2019 - suspended development*
- A minimum Go compiler in Go that aims to do self-hosting. Its grammar is based on the [official specification](https://golang.org/ref/spec), but it only supports parts of them. It can compile the Go file which calculate fibonacci numbers. It is now suspended development.

**[mdtohtml](https://github.com/d0iasm/mdtohtml)** | *March, October 2019*
- A HTML generator from a Markdown file implemented in Go. The unique feature of this tool is that default result file includes CSS, but you can skip to include CSS by passing -nocss parameter.

**[haribote OS](https://github.com/d0iasm/haribote-os)** | *April 2018 - December 2018*
- Haribote OS based on x86, 32-bit in C. This code refers to the Japanese book "[OS自作入門](https://www.amazon.co.jp/dp/4839919844/ref=cm_sw_em_r_mt_dp_U_DYAPDbA18822Y)".

**[Life Game](https://github.com/d0iasm/life-game.cpp)** | *September 2018*
- Life game with WebAssembly compiled from C++. The [site](https://d0iasm.github.io/life-game.cpp/) is hosted on GitHub pages.

**[Graphy](https://github.com/d0iasm/graphy)** | *June 2017 - February 2018*
- Slack BOT to generate an image based on text messages in Python. It's hosted on DigitalOcean and AWS S3.The goal of Graphy is to help us decide whether or not to read all our unread messages, and to let us get an idea of their content before reading them. This bot creates an image from the text by using dependency analysis. In addition, the nodes of important keywords are emphasized by color and size.

**[History Search](https://github.com/d0iasm/history-search)** | *August 2017*
- Chrome browser extension in HTML, CSS, and JavaScript which can search the page you want to see again from your browser history. You can use it in [the chrome webstore](https://chrome.google.com/webstore/detail/history-search/bbmclnpfclejopgaicmhgpocicpijodj).

**[Photo Editor Bot](https://github.com/d0iasm/photo-editor-bot)** | *February 2017*
- This is my first big project in PHP (Slim)! It was hosted on Heroku. I am interested in how to work efficiently. Thus many tasks have to work automatically as much as possible. In this age that people upload many pictures to the SNS, photo editing is necessary but it would waste of the time. This BOT works  based on LINE, so people can invite it to thier chat group and the BOT automatically edits photos.

## OSS Contributions
**Wasmer**
- wasmerio/wasmer
  - [Adapt backend usage depending on wasm file executed](https://github.com/wasmerio/wasmer/pull/1004)
- wasmerio/go-ext-wasm
  - [Upgrade the cli package version to v2](https://github.com/wasmerio/go-ext-wasm/pull/110)

**coreboot**
- [Commits list in Gerrit](https://review.coreboot.org/q/owner:d0iasm)

**Chromium**
- [Commits list in Gerrit](https://chromium-review.googlesource.com/q/owner:asamidoi)

**rustwasm/book**
- [Update the template source code](https://github.com/rustwasm/book/pull/194)

**os/Slacker**
- [Update example files from python 2 to python 3](https://github.com/os/slacker/pull/129)
- [Add the new method of chat.getPermalink](https://github.com/os/slacker/pull/130)

## Talks
- "WebAssembly outside of the browser" at [カーネル/VM探検隊@関西 10回目](https://connpass.com/event/161201/) | *February 2020*
  - [slide (jp)](https://speakerdeck.com/d0iasm/webassembly-outside-of-the-browser)
  - [YouTube (jp)](https://youtu.be/J-pF4fg3r04)
- "FFI: RustとWebAssemblyとJavaScriptと" at [Rust LT #8](https://rust.connpass.com/event/160225/) | *January 2020*
  - [slide (jp)](https://speakerdeck.com/d0iasm/ffi-rusttowebassemblytojavascriptto)
- "Hello World in coreboot" at [第15回 カーネル/VM探検隊](https://kernelvm15.peatix.com/) | *July 2019*
  - [slide (jp)](https://speakerdeck.com/d0iasm/hello-world-in-coreboot-9bdb3d94-d092-4a60-bb19-360f7fb777df)
  - [YouTube (jp)](https://youtu.be/RNzoBRffVdE)

## Awards
- STEP Coding Camp | *March 2018*
  - Won an award in 3 days Hackathon at Google by implementing a search algorithm in a team.
- Life Style Hackathon | *February 2017*
  - Won an award in 3 days Hackathon in NTT DATA by creating a prototype of a new business.
- Online code challenge of VOYAGE GROUP | *September 2016*
  - Joined this competitive programming content by using Java. Ranked as one of the top 10.
- OthloHack 2016 | *August 2016*
  - Won awards in 1 day Hackathon by suggesting new applications and create a prototype.

## Miscellaneous
- Study abroad at Goldsmiths, University of London | *October 2019*
- Orginize a game jam at Nagoya University | *August 2019*
  - [event page](https://nu-info.connpass.com/event/139070/)
- Security next camp 2019 | *August 2019*
- Crisis management contest 12th, pass the first qualifying round | *April 2019*
- ELS language school in Cupertino, California | *January 2019 - February 2019*
- Cisco Live! 2018 in Orlando | *June 2018*
  - [Interview at Cisco Live! 2018 in Orlando](https://www.cisco.com/c/m/ja_jp/about/security-scholarship/security-scholarship-clus2018.html)
  - It is the one of the program which is the cyber security scholarship at Cisco | *March 2018*
- STEP development course at Google | *May 2017 - August 2017*
  - Took CS classes at Google Japan. There were homework almost every week and I submitted all of them in Python and Go.

<br />
---

This page generated by the project [d0iasm/mdtohtml](https://github.com/d0iasm/mdtohtml) written in Go which converts a markdown file to a HTML file including CSS design.
