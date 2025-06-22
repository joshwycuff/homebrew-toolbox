class Typescore < Formula
  desc "Simple tool to score the typing difficulty of text"
  homepage "https://github.com/joshwycuff/homebrew-toolbox"
  url "https://github.com/joshwycuff/homebrew-toolbox/releases/download/typescore-2c93a9d538e6fccd455dc908a82e3bd876ea99ee/typescore.tar.gz"
  version "2c93a9d538e6fccd455dc908a82e3bd876ea99ee"

  def install
    bin.install "typescore"
    man1.install Dir["*.1"]
    bash_completion.install "typescore.bash" => "typescore"
    zsh_completion.install "typescore.zsh" => "_typescore"
    fish_completion.install "typescore.fish"
  end
end
