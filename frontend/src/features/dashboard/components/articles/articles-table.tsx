import Link from "next/link";
import {
  Table,
  TableTr,
  TableTd,
  TableTh,
  TableThead,
  TableTbody,
  ActionIcon,
  ActionIconGroup,
  Tooltip,
  Group,
  Button,
  Badge,
  rem,
} from "@mantine/core";
import {ArticlesPagination} from "./articles-table-pagination";
import {ArticleDeleteButton} from "./article-delete-button";
import {IconEye, IconPencil, IconFilePlus} from "@tabler/icons-react";
import {fetchAllArticles} from "@/dal/articles";
import {dateFromNow} from "@/lib/date-and-time";
import {APP_PATHS} from "@/lib/app-paths";

type Props = {
  page: number | string;
};

export async function ArticlesTable({page}: Props) {
  const articlesResponse = await fetchAllArticles({
    params: {
      page: page,
    },
  });
  const articles = articlesResponse.items;
  const {total_pages, current_page} = articlesResponse.pagination;

  const tableActions = [
    {
      tooltipLabel: "بازدید کردن مقاله",
      Icon: IconEye,
      color: "blue",
      href: (uuid: string) => APP_PATHS.articles.detail(uuid),
      disabled: (published: boolean) => published,
    },
    {
      tooltipLabel: "ویرایش کردن مقاله",
      Icon: IconPencil,
      color: "blue",
      href: (uuid: string) => APP_PATHS.dashboard.articles.edit(uuid),
      disabled: () => false,
    },
  ];

  return (
    <>
      <Group justify="flex-end">
        <Button
          variant="light"
          component={Link}
          leftSection={<IconFilePlus />}
          href={APP_PATHS.dashboard.articles.new}
        >
          مقاله جدید
        </Button>
      </Group>
      <Table verticalSpacing={"sm"} striped withRowBorders>
        <TableThead>
          <TableTr>
            <TableTh>#</TableTh>
            <TableTh>عنوان</TableTh>
            <TableTh>تاریخ انتشار</TableTh>
            <TableTh>عملیات</TableTh>
          </TableTr>
        </TableThead>
        <TableTbody>
          {articles.length === 0 && (
            <TableTr>
              <TableTd colSpan={4} ta={"center"}>
                مقاله های وجود ندارد
              </TableTd>
            </TableTr>
          )}
          {articles.map((article: any, index: number) => {
            const isPublished = new Date(article.published_at).getDate() !== 1;

            return (
              <TableTr key={article.uuid}>
                <TableTd>{index + 1}</TableTd>
                <TableTd>{article.title}</TableTd>
                <TableTd>
                  {isPublished ? (
                    dateFromNow(article.published_at)
                  ) : (
                    <Badge color="yellow" variant="light">
                      منتشر نشده
                    </Badge>
                  )}
                </TableTd>
                <TableTd>
                  <ActionIconGroup>
                    {tableActions.map(
                      ({Icon, tooltipLabel, color, href, disabled}) => {
                        return (
                          <Tooltip
                            key={tooltipLabel}
                            label={tooltipLabel}
                            withArrow
                          >
                            <ActionIcon
                              component={Link}
                              variant="light"
                              size="lg"
                              color={color}
                              href={href(article.uuid)}
                              disabled={disabled(isPublished === false)}
                              aria-label={tooltipLabel}
                            >
                              <Icon style={{width: rem(20)}} stroke={1.5} />
                            </ActionIcon>
                          </Tooltip>
                        );
                      },
                    )}
                    <ArticleDeleteButton
                      articleID={article.uuid}
                      articleTitle={article.title}
                    />
                  </ActionIconGroup>
                </TableTd>
              </TableTr>
            );
          })}
        </TableTbody>
      </Table>
      {articles.length >= 1 && (
        <Group mt="md" mb={"lg"} justify="flex-end">
          <ArticlesPagination total={total_pages} current={current_page} />
        </Group>
      )}
    </>
  );
}
