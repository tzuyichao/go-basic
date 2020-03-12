    
        String processedRawText;
        if(rawTextPreProcessor != null) {
            processedRawText = rawTextPreProcessor.apply(rawText);
        } else {
            processedRawText = rawText;
        }
        QLangErrorListener surplusErrorListener = new QLangErrorListener();
        CharStream input = CharStreams.fromString(processedRawText);
        qlang.QLangLexer lexer = new qlang.QLangLexer(input);
        lexer.removeErrorListeners();
        lexer.addErrorListener(surplusErrorListener);

        CommonTokenStream tokens = new CommonTokenStream(lexer);
        qlang.QLangParser parser = new qlang.QLangParser(tokens);
        parser.removeErrorListeners();
        parser.addErrorListener(surplusErrorListener);

        ParseTree parseTree = parser.prog();

        if (!surplusErrorListener.isSyntaxError()) {
            ParseTreeWalker walker = new ParseTreeWalker();
            walker.walk((ParseTreeListener) evaluator, parseTree);
            QueryString result = evaluator.getResult();
            logger.debug(result.toString());
            return result;
        } else {
            logger.error("Syntax Error");
            logger.error(surplusErrorListener.getSyntaxErrorMessages().toString());
            return new QueryString(rawText);
        }