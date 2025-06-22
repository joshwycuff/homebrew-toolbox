class Typescore < Formula
  desc "Simple tool to score the typing difficulty of text"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/typescore-654f46bc058ea34cb4bf1373b439b902cf699b1f/typescore.tar.gz"
  version "654f46bc058ea34cb4bf1373b439b902cf699b1f"

  def install
    bin.install "typescore"
    man1.install Dir["*.1"]
    bash_completion.install "typescore.bash" => "typescore"
    zsh_completion.install "typescore.zsh" => "_typescore"
    fish_completion.install "typescore.fish"
  end
end
