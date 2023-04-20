#!/bin/sh

sqlcmd -S uswapserveractual.database.windows.net -U LeSwapper -P Uswapcen$ -d uswapDatabase -i ./CreateTestData.sql