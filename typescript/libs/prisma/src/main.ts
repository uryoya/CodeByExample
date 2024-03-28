import { PrismaClient } from "@prisma/client";

const prisma = new PrismaClient({
  log: ["query"],
});

async function main() {
  // const imagesFromMediaTable = await prisma.media.findMany({
  //   include: {
  //     image: true,
  //   },
  //   where: {
  //     type: "image",
  //   },
  //   orderBy: {
  //     title: "asc",
  //   },
  //   take: 10,
  // });

  // console.log(imagesFromMediaTable.length);

  const imagesFromImageTable = await prisma.image.findMany({
    include: {
      media: true,
    },
    orderBy: {
      media: {
        title: "asc",
      },
    },
    take: 10,
  });

  console.log(imagesFromImageTable.length);
}

main();
