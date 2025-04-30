#syntax=harbor.leops.local/library/docker/dockerfile:1

# ---- 编译环境 ----
FROM harbor.leops.local/common/tools/golang:1.24 AS builder

ARG APP_ENV=test \
  APP=undefine \
  GIT_BRANCH= \
  GIT_COMMIT_ID=

ENV APP_ENV=$APP_ENV \
  APP=$APP \
  GIT_BRANCH=$GIT_BRANCH \
  GIT_COMMIT_ID=$GIT_COMMIT_ID

WORKDIR /app_build

# 编译
COPY . .
RUN --mount=type=cache,id=gomod,target=/go/pkg/mod \
  --mount=type=cache,id=gobuild,target=/root/.cache/go-build \
  go build  -tags 'osusergo,netgo' \
  -ldflags "-X main.branch=$GIT_BRANCH -X main.commit=$GIT_COMMIT_ID" \
  -v -o bin/${APP} *.go \
  && cp -rf etc bin/etc \
  && chown 999.999 -R bin

#
# ---- 运行环境 ----
FROM harbor.leops.local/common/runtime/golang:debian11 AS running

ARG APP_ENV=test \
  APP=undefine \
  GIT_BRANCH= \
  GIT_COMMIT_ID=

ENV APP_ENV=$APP_ENV \
  APP=$APP \
  GIT_BRANCH=$GIT_BRANCH \
  GIT_COMMIT_ID=$GIT_COMMIT_ID

WORKDIR /app

COPY --from=builder --link /app_build/bin /app/

CMD ["bash", "-c", "exec /app/${APP} -f /app/etc/app_${APP_ENV}.yaml"]