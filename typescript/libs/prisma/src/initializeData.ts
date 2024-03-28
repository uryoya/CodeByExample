import { Prisma, PrismaClient } from "@prisma/client";
import { randomUUID } from "crypto";

const prisma = new PrismaClient();

const range = (start: number, end: number) => {
  let arr = [];
  for (let i = start; i <= end; i++) {
    arr.push(i);
  }
  return arr;
};

const images: Prisma.MediaCreateArgs[] = range(1, 10000).map((i) => {
  const id = randomUUID();
  return {
    data: {
      id,
      title: `Image ${i}`,
      type: "image",
      url: `https://example.com/image/${id}.jpg`,
      image: {
        create: {
          height: 100,
          width: 100,
        },
      },
    },
  };
});

prisma.$transaction(images.map((args) => prisma.media.create(args)));

const audios: Prisma.MediaCreateArgs[] = range(1, 10000).map((i) => {
  const id = randomUUID();
  return {
    data: {
      id,
      title: `Audio ${i}`,
      type: "audio",
      url: `https://example.com/audio/${id}.mp3`,
      audio: {
        create: {
          duration: 100,
        },
      },
    },
  };
});

prisma.$transaction(audios.map((args) => prisma.media.create(args)));

const videos: Prisma.MediaCreateArgs[] = range(1, 10000).map((i) => {
  const id = randomUUID();
  return {
    data: {
      id,
      title: `Video ${i}`,
      type: "video",
      url: `https://example.com/video/${id}.mp4`,
      video: {
        create: {
          duration: 300,
          height: 100,
          width: 100,
        },
      },
    },
  };
});

prisma.$transaction(videos.map((args) => prisma.media.create(args)));
